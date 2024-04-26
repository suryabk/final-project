package repository

import (
	"database/sql"
	"final-project/structs"
	"time"
)

func GetAllTask(db *sql.DB) (results []structs.Task, err error) {
	sql := "SELECT * FROM tasks"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var task = structs.Task{}

		err := rows.Scan(&task.TaskID, &task.TaskName, &task.Description, &task.Priority, &task.Deadline, &task.StatusID, &task.AssignedTo, &task.ProjectID, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, task)
	}

	return
}

func InsertTask(db *sql.DB, task structs.Task) (err error) {
	sql := "INSERT INTO tasks (task_id, task_name, description, priority, deadline, status_id, assigned_to, project_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	errs := db.QueryRow(sql, task.TaskName, task.Description, task.Priority, task.Deadline, task.StatusID, task.AssignedTo, task.ProjectID, time.Now(), time.Now())

	return errs.Err()
}

func UpdateTask(db *sql.DB, task structs.Task) (err error) {
	sql := "UPDATE tasks SET task_name = $1, description = $2, priority = $3, deadline = $4, status_id = $5, assigned_to = $6, project_id = $7, updated_at = $8  WHERE id = $9"
	errs := db.QueryRow(sql, task.TaskName, task.Description, task.Priority, task.Deadline, task.StatusID, task.AssignedTo, task.ProjectID, time.Now(), task.TaskID)

	return errs.Err()
}

func DeleteTask(db *sql.DB, task structs.Task) (err error) {
	sql := "DELETE FROM tasks WHERE id = $1"
	errs := db.QueryRow(sql, task.TaskID)

	return errs.Err()
}
