package main

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/storage"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func mustInitDB(dataSource string) *sql.DB {
	const name = "postgres-log"
	sql.Register(name, drv{})
	db, err := sql.Open(name, dataSource)
	if err != nil {
		panic(err)
	}
	return db
}

func mustExec(db *sql.DB, query string, params ...interface{}) {
	if _, err := db.Exec(query, params...); err != nil {
		panic(err)
	}
}

func logQuery(query string, args interface{}) {
	as := fmt.Sprint(args)
	if len(as) > 100 {
		as = as[:100] + "..."
	}
	log.Printf("Query: %v: %s", as, query)
}

func trace(s string) (string, time.Time) {
	return s, time.Now()
}

func un(s string, start time.Time) {
	log.Printf("%s: %s", time.Since(start), s)
}

type drv struct{}

func (d drv) Open(name string) (driver.Conn, error) {
	c, err := pq.Open(name)
	if *flagInit {
		// Verbose logging for local dev.
		c = &conn{c}
	}
	return c, err
}

// conn implements a logging driver.Conn that logs queries.
type conn struct {
	driver.Conn
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	logQuery(query, "[prepare]")
	return c.Conn.Prepare(query)
}

func (c *conn) Begin() (driver.Tx, error) {
	logQuery("Begin()", nil)
	return c.Conn.Begin()
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	logQuery("BeginTx()", nil)
	return c.Conn.(driver.ConnBeginTx).BeginTx(ctx, opts)
}

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	logQuery(query, args)
	c.ExplainNamed(query, args)
	un(trace(query))
	return c.Conn.(driver.QueryerContext).QueryContext(ctx, query, args)
}

func (c *conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	logQuery(query, args)
	c.Explain(query, args)
	un(trace(query))
	return c.Conn.(driver.Queryer).Query(query, args)
}

func (c *conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	logQuery(query, args)
	return c.Conn.(driver.Execer).Exec(query, args)
}

func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (sql.Result, error) {
	logQuery(query, args)
	return c.Conn.(driver.ExecerContext).ExecContext(ctx, query, args)
}

func (c *conn) ExplainNamed(query string, args []driver.NamedValue) {
	values := make([]driver.Value, len(args))
	for i, v := range args {
		values[i] = v.Value
	}
	c.Explain(query, values)
}

func (c *conn) Explain(query string, args []driver.Value) {
	fmt.Println("EXPLAIN", query, args)
	rows, err := c.Conn.(driver.Queryer).Query("EXPLAIN  "+query, args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var level int64
	var typ, field, description string
	values := []driver.Value{level, typ, field, description}
	for {
		if err := rows.Next(values); err == io.EOF {
			return
		} else if err != nil {
			fmt.Println(err)
		}
		fmt.Println("EXPLAIN", values)
	}
}

func makeValues(numArgs, start int) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	for i := 0; i < numArgs; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "$%d", i+start)
	}
	buf.WriteString(")")
	return buf.String()
}

func (h *hotsContext) Import(bucket string, max int) error {
	mustExec(h.db, `CREATE DATABASE IF NOT EXISTS `+dbName)
	count := max/perFile + 1
	args := make([]interface{}, count)
	for i := 0; i < count; i++ {
		args[i] = fmt.Sprintf("gs://%s/game/"+configBase, bucket, i*perFile)
	}
	if _, err := h.db.Exec(fmt.Sprintf(`
			IMPORT TABLE games (
				id INT PRIMARY KEY,
				mode INT,
				time TIMESTAMP,
				map INT,
				length INT,
				build INT,
				region INT,

				bans INT[]
			) CSV DATA %s
		`, makeValues(count, 1)),
		args...); err != nil {
		return errors.Wrap(err, "import games")
	}
	for i := 0; i < count; i++ {
		args[i] = fmt.Sprintf("gs://%s/player/"+configBase, bucket, i*perFile)
	}
	if _, err := h.db.Exec(fmt.Sprintf(`
			IMPORT TABLE players (
				game INT,
				mode INT,
				time TIMESTAMP,
				map INT,
				length INT,
				build INT,
				region INT,

				hero INT,
				hero_level INT,
				team INT,
				winner BOOL,
				blizzid INT,
				skill INT,
				battletag STRING COLLATE en_u_ks_level1,
				talents INT[],
				data JSONB,

				PRIMARY KEY (game, blizzid),

				INDEX (region, blizzid, time DESC),
				INDEX (region, battletag),

				INDEX (build, map, mode, hero_level) STORING (hero, winner),
				INDEX (build, map, hero_level) STORING (hero, winner),
				INDEX (build, mode, hero_level) STORING (hero, winner),
				INDEX (build, hero_level) STORING (hero, winner),
				INDEX (build, hero) STORING (winner, hero_level, length, map, mode),
				INDEX (build, hero, map, mode, hero_level) STORING (winner, talents),
				INDEX (build, hero, map, hero_level) STORING (winner, talents),
				INDEX (build, hero, mode, hero_level) STORING (winner, talents),
				INDEX (build, hero, hero_level) STORING (winner, talents)
			) CSV DATA %s
		`, makeValues(count, 1)),
		args...); err != nil {
		return errors.Wrap(err, "import games")
	}
	return nil
}

func (h *hotsContext) syncConfig(bucket string) error {
	ctx := context.Background()
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "new client")
	}
	handle := cl.Bucket(bucket)
	config, err := getConfig(ctx, handle)
	if err != nil {
		return errors.Wrap(err, "get config")
	}
	b, err := json.Marshal(&config)
	if err != nil {
		return errors.Wrap(err, "encode config")
	}
	if _, err := h.db.Exec(`UPSERT INTO config (key, s) VALUES ($1, $2)`, cacheConfig, b); err != nil {
		return errors.Wrap(err, "upsert config")
	}
	return nil
}

const cacheConfig = "cacheconfig"
