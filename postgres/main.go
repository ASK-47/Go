package main

import (
	"context"
	"fmt"
	"postgres/freature_postgres/simple_connection"
	"postgres/freature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.Create_connetion(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Eat the cat"
			task.Description = "Take a cat"
			task.Completed = true
			now := time.Now()
			task.Completed_at = &now

			if err := simple_sql.Update(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}

	/*if err := simple_sql.InsertRow(
		ctx,
		conn,
		"Wqlk along",
		"Wqlk along far away",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}*/

	/*if err := simple_sql.UpdatetRow(ctx, conn); err != nil {
		panic(err)
	}*/

	/*if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}*/

	//fmt.Println("=====================================================")
	//pp.Println(tasks)
	//fmt.Println(tasks)

	fmt.Println("Table was successfully created")

	/*val := os.Getenv("phone_number")
	if val != "" {
		fmt.Println("val=", val)

	} else {
		fmt.Println("variable is not initiated")
	}*/
}
