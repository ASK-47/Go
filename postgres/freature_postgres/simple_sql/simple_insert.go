package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
	/*title string,
	description string,
	completed bool,
	created_at time.Time,*/
) error {
	/*sqlQery := `
		INSERT INTO tasks (title ,description, completed, created_at)
		VALUES ('Breakfast', 'To eat something', FALSE, '2026-04-19 18:18:18');
	`
	_, err := conn.Exec(ctx, sqlQery)*/

	sqlQery := `
		INSERT INTO tasks (title, description, completed, created_at)
		VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(ctx, sqlQery, task.Title, task.Description, task.Completed, task.Completed_at)
	return err
}
