package repository

import (
	"database/sql"
	"final-project/structs"
)

func GetAllTaskStatus(db *sql.DB) (results []structs.TaskStatus, err error) {
	sql := "SELECT * FROM task_status"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var taskStatus = structs.TaskStatus{}

		err := rows.Scan(&taskStatus.StatusID, taskStatus.StatusName)
		if err != nil {
			panic(err)
		}

		results = append(results, taskStatus)
	}

	return
}

func InsertTaskStatus(db *sql.DB, status structs.TaskStatus) (err error) {
	sql := "INSERT INTO task_status (status_id, status_name) VALUES ($1, $2)"
	errs := db.QueryRow(sql, status.StatusID, status.StatusName)

	return errs.Err()
}

func UpdateTaskStatus(db *sql.DB, status structs.TaskStatus) (err error) {
	sql := "UPDATE task_status SET status_name = $1 WHERE id = $2"
	errs := db.QueryRow(sql, status.StatusName, status.StatusName)

	return errs.Err()
}

func DeleteTaskStatus(db *sql.DB, status structs.TaskStatus) (err error) {
	sql := "DELETE FROM task_status WHERE id = $1"
	errs := db.QueryRow(sql, status.StatusID)

	return errs.Err()
}
