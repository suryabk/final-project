package repository

import (
	"database/sql"
	"final-project/structs"
	"time"
)

func GetAllProject(db *sql.DB) (results []structs.Project, err error) {
	sql := "SELECT * FROM projects"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var project = structs.Project{}

		err := rows.Scan(&project.ProjectID, &project.ProjectName, &project.Description, &project.Budget, &project.CreatedBy, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, project)
	}

	return
}

func InsertProject(db *sql.DB, project structs.Project) (err error) {
	sql := "INSERT INTO projects (project_name, description, budget, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	errs := db.QueryRow(sql, project.ProjectName, project.Description, project.Budget, project.CreatedBy, time.Now(), time.Now())

	return errs.Err()
}

func UpdateProject(db *sql.DB, project structs.Project) (err error) {
	sql := "UPDATE projects SET project_name = $1, description = $2, budget = $3, created_by = $4, updated_at = $5 WHERE id = $6"
	errs := db.QueryRow(sql, project.ProjectName, project.Description, project.Budget, project.CreatedBy, time.Now(), project.ProjectID)

	return errs.Err()
}

func DeleteProject(db *sql.DB, p structs.Project) (err error) {
	sql := "DELETE FROM projects WHERE id = $1"
	errs := db.QueryRow(sql, p.ProjectID)

	return errs.Err()
}
