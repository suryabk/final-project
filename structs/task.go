package structs

import "time"

type Task struct {
	TaskID      int       `json:"task_id"`
	TaskName    string    `json:"task_name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Deadline    time.Time `json:"deadline"`
	StatusID    int       `json:"status_id"`
	AssignedTo  int       `json:"assigned_to"`
	ProjectID   int       `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InputTask struct {
	TaskName    string    `json:"task_name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Deadline    time.Time `json:"deadline"`
	AssignedTo  int       `json:"assigned_to"`
	ProjectID   int       `json:"project_id"`
}
