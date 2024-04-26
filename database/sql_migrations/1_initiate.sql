CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE task_status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(20) NOT NULL UNIQUE,
);

INSERT INTO task_status (status_name) VALUES
    ('Open'),
    ('In Progress'),
    ('Closed'),

CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    project_name VARCHAR(100) NOT NULL,
    description TEXT,
    budget DECIMAL(10, 2),
    deadline DATE,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    task_name VARCHAR(100) NOT NULL,
    description TEXT,
    priority INT,
    deadline DATE,
    status_id INTEGER DEFAULT 1 REFERENCES task_status(status_id), -- Nilai default 'Open'
    assigned_to INTEGER REFERENCES users(user_id),
    project_id INTEGER REFERENCES projects(project_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);