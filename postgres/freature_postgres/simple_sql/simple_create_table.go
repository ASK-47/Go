package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQery := `
		CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		comleted_at TIMESTAMP, 
		UNIQUE (title)
	);
	`
	_, err := conn.Exec(ctx, sqlQery)
	/*if err != nil {
		return err
	}
	return nil*/
	return err
}
