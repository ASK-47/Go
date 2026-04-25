package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQery := `
		SELECT id, title, description, completed, created_at, comleted_at
		FROM tasks
		ORDER BY id ASC
	`
	rows, err := conn.Query(ctx, sqlQery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel
		/*var id int
		var title string
		var description string
		var completed bool
		var created_at time.Time
		var comleted_at *time.Time*/
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.Created_at, &task.Completed_at)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		//fmt.Println(id, title, description, completed, created_at, comleted_at)
		print(task)

	}
	return tasks, nil
}

func print(task TaskModel) {
	fmt.Println("=====================================================")
	fmt.Println("id:\t", task.ID)
	fmt.Println("title:\t", task.Title)
	fmt.Println("description:\t", task.Description)
	fmt.Println("completed:\t", task.Completed)
	fmt.Println("created_at:\t", task.Created_at)
	fmt.Println("comleted_at:\t", task.Completed_at)
}
