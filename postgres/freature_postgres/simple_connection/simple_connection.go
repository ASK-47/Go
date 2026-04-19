package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Create_connetion(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:Huw58m@localhost:5432/postgres")
}
