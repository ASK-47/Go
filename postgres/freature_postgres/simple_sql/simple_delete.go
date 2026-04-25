package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, taskIDs []int) error {
	/*sqlQery := `
		DELETE FROM tasks
		WHERE ID BETWEEN 6 AND 15;
	`*/
	sqlQery := `
		DELETE FROM tasks		
		WHERE id =ANY($1);
	`
	_, err := conn.Exec(ctx, sqlQery, taskIDs)
	return err
}
