package structs

import "time"

type User struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Project struct {
	ProjectID   int       `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Description string    `json:"description"`
	Budget      float64   `json:"budget"`
	Deadline    time.Time `json:"deadline"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InputProject struct {
	ProjectName string    `json:"project_name"`
	Description string    `json:"description"`
	Budget      float64   `json:"budget"`
	Deadline    time.Time `json:"deadline"`
}

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

type TaskStatus struct {
	StatusID   int    `json:"status_id"`
	StatusName string `json:"status_name"`
}
