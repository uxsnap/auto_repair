package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Client struct {
	dbc *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Client, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Errorf("failed to ping db: %v", err)
	}

	return &Client{
		dbc: pool,
	}, nil
}

func (c *Client) GetDb() *pgxpool.Pool {
	return c.dbc
}

func (c *Client) Close() error {
	if c.dbc != nil {
		c.dbc.Close()
	}

	return nil
}
