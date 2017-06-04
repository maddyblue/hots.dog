package main

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
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

func logQuery(query string, args ...interface{}) {
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
	return c.Conn.(driver.QueryerContext).QueryContext(ctx, query, args)
}

func (c *conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	logQuery(query, args)
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

func makeValues(numArgs, argCount int) string {
	var buf bytes.Buffer
	for i := 0; i < argCount; i++ {
		m := i % numArgs
		if m == 0 {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString("(")
		}
		fmt.Fprintf(&buf, "$%d", i+1)
		if m == numArgs-1 {
			buf.WriteString(")")
		} else {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}
