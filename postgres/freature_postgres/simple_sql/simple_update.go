package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdatetRow(ctx context.Context, conn *pgx.Conn) error {
	/*sqlQery := `
		UPDATE tasks
		SET completed = true
		WHERE ID=4
	`*/

	sqlQery := `
		UPDATE tasks
		SET description = 'AAAAAAA'
		WHERE completed=false
	`
	_, err := conn.Exec(ctx, sqlQery)
	return err
}

/*func UpdateTitle(ctx context.Context, conn *pgx.Conn, id int, newTitle string) error {
	sqlQery := `
		UPDATE tasks
		SET title =$1
		WHERE id=$2
	`
	_, err := conn.Exec(ctx, sqlQery, newTitle)
	return err
}*/

func Update(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQery := `
		UPDATE tasks
		SET title =$1, description =$2, completed=$3, created_at =$4, comleted_at=$5 		
		WHERE id=$6
	`
	_, err := conn.Exec(ctx, sqlQery, task.Title, task.Description, task.Completed, task.Created_at, task.Completed_at, task.ID)
	return err
}
