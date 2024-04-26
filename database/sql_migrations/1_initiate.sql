-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE task_status (
    status_id SERIAL PRIMARY KEY NOT NULL,
    status_name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY NOT NULL,
    project_name VARCHAR(100) NOT NULL,
    description TEXT,
    budget DECIMAL(10, 2),
    -- deadline DATE,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY NOT NULL,
    task_name VARCHAR(100) NOT NULL,
    description TEXT,
    priority VARCHAR(20),
    -- deadline DATE,
    status_id INTEGER REFERENCES task_status(status_id), 
    assigned_to INTEGER REFERENCES users(user_id),
    project_id INTEGER REFERENCES projects(project_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +migrate StatementEnd