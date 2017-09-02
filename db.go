package main

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"

	"github.com/lib/pq"
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

func logQuery(query string, args interface{}) {
	log.Printf("Query: %v: %s", args, query)
}

type drv struct{}

func (d drv) Open(name string) (driver.Conn, error) {
	c, err := pq.Open(name)
	return &conn{c}, err
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
	return c.Conn.(driver.QueryerContext).QueryContext(ctx, query, args)
}

func (c *conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	logQuery(query, args)
	c.Explain(query, args)
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

func makeValues(numArgs int) string {
	var buf bytes.Buffer
	buf.WriteString("(")
	for i := 1; i <= numArgs; i++ {
		if i > 1 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "$%d", i)
	}
	buf.WriteString(")")
	return buf.String()
}
