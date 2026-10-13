package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,

		UNIQUE(title)
	);`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);`

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt)
	return err
}

func UpdateRow(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE
	WHERE id = $1;`

	_, err := conn.Exec(ctx, sqlQuery, id)
	return err
}

func DeleteRow(ctx context.Context, conn *pgx.Conn, taskIDs []int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = ANY ($1);`

	_, err := conn.Exec(ctx, sqlQuery, taskIDs)
	return err
}

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks;`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []TaskModel{}
	for rows.Next() {
		task := TaskModel{}

		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
