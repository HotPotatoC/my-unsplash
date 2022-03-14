package clients

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresClient struct {
	conn *pgxpool.Pool
}

// NewPostgreSQLClient creates a new postgresql database instance
func NewPostgreSQLClient(ctx context.Context, connString string) (*PostgresClient, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &PostgresClient{
		conn: conn,
	}, nil
}

func (c *PostgresClient) Exec(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	result, err := c.conn.Exec(ctx, sql, args...)
	return result.RowsAffected(), err
}

func (c *PostgresClient) Query(ctx context.Context, sql string, args ...interface{}) (*ClientRows, error) {
	rows, err := c.conn.Query(ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (c *PostgresClient) QueryRow(ctx context.Context,sql string, args ...interface{}) *ClientRow {
	row := c.conn.QueryRow(ctx, sql, args...)
	return newDatabaseRow(row)
}

type ClientRow struct {
	row pgx.Row
}

func newDatabaseRow(row pgx.Row) *ClientRow {
	return &ClientRow{
		row: row,
	}
}

func (c *ClientRow) Scan(dest ...interface{}) error {
	return c.row.Scan(dest...)
}

type ClientRows struct {
	rows pgx.Rows
}

func newDatabaseRows(rows pgx.Rows) *ClientRows {
	return &ClientRows{
		rows: rows,
	}
}

func (c *ClientRows) Scan(dest ...interface{}) error {
	return c.rows.Scan(dest...)
}

func (c *ClientRows) Next() bool {
	return c.rows.Next()
}

func (c *ClientRows) Close() {
	c.rows.Close()
}

func (c *ClientRows) Err() error {
	return c.rows.Err()
}

type Transaction struct {
	ctx context.Context
	tx  pgx.Tx
}

func NewDatabaseTransaction(ctx context.Context, tx pgx.Tx) *Transaction {
	return &Transaction{
		ctx: ctx,
		tx:  tx,
	}
}

func (c *Transaction) Commit() error {
	return c.tx.Commit(c.ctx)
}

func (c *Transaction) Rollback() error {
	return c.tx.Rollback(c.ctx)
}

func (c *Transaction) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := c.tx.Exec(c.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (c *Transaction) Query(sql string, args ...interface{}) (*ClientRows, error) {
	rows, err := c.tx.Query(c.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (c *Transaction) QueryRow(sql string, args ...interface{}) *ClientRow {
	row := c.tx.QueryRow(c.ctx, sql, args...)
	return newDatabaseRow(row)
}
