package repository

import (
	"database/sql"
	"final-project/structs"
	"time"
)

func GetAllProject(db *sql.DB) (results []structs.Project, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var project = structs.Project{}

		err := rows.Scan(&project.ProjectID, &project.ProjectName, &project.Description, &project.Budget, &project.Deadline, &project.CreatedBy, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, project)
	}

	return
}

func InsertProject(db *sql.DB, project structs.Project) (err error) {
	sql := "INSERT INTO projects (project_id, project_name, description, budget, deadline, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, project.ProjectName, project.Description, project.Budget, project.Deadline, project.CreatedBy, time.Now(), time.Now())

	return errs.Err()
}

func UpdateProject(db *sql.DB, project structs.Project) (err error) {
	sql := "UPDATE projects SET project_name = $1, description = $2, budget = $3, deadline = $4, created_by = $5, updated_at = $6 WHERE id = $7"
	errs := db.QueryRow(sql, project.ProjectName, project.Description, project.Budget, project.Deadline, project.CreatedBy, time.Now(), project.ProjectID)

	return errs.Err()
}

func DeleteProject(db *sql.DB, p structs.Project) (err error) {
	sql := "DELETE FROM projects WHERE id = $1"
	errs := db.QueryRow(sql, p.ProjectID)

	return errs.Err()
}
