package structs

import "time"

type Project struct {
	ProjectID   int     `json:"project_id"`
	ProjectName string  `json:"project_name"`
	Description string  `json:"description"`
	Budget      float64 `json:"budget"`
	// Deadline    time.Time `json:"deadline"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputProject struct {
	ProjectName string  `json:"project_name"`
	Description string  `json:"description"`
	Budget      float64 `json:"budget"`
	// Deadline    string  `json:"deadline"`
}
