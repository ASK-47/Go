package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Create_connetion(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	//return pgx.Connect(ctx, "postgres://postgres:Huw58m@localhost:5432/postgres")
	return pgx.Connect(ctx, connString)

	//need in terminal: export conn_string=postgres://postgres:Huw58m@localhost:5432/postgres
}
