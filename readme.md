# Project Management API

## Introduction

This document provides an overview of the routes available in the Project Management API.

## Routes List

### /register

- **Method**: POST
- **Description**: Register a new user
 - **Parameters**:
    - `email` (string): email must unique
    - `password` (string): password account

### /login

- **Method**: POST
- **Description**: Log in with existing user credentials
- **Parameters**:
    - `email` (string): email must unique
    - `password` (string): password account

### /task-status

- **Parameters for POST and PUT**:
    - `status_name` (string): Name of the task status

- **Method**: GET
- **Description**: Get all task statuses

- **Method**: POST
- **Description**: Insert a new task status

- **Method**: PUT
- **Description**: Update a task status by ID

- **Method**: DELETE
- **Description**: Delete a task status by ID

### /task

- **Parameters for POST and PUT**:
    - `task_name` (string): Name of the task
    - `description` (string): Description of the task
    - `priority` (string): Priority level of the task
    - `status_id` (number): ID of the task status based on **task_status(status_id)**
    - `assigned_to` (number): ID of the user assigned to the task based on**task users(user_id)**
    - `project_id` (number): ID of the project the task belongs to based on **projects(project_id)**
- **Method**: GET
- **Description**: Get all tasks

- **Method**: POST
- **Description**: Insert a new task

- **Method**: PUT
- **Description**: Update a task by ID

- **Method**: DELETE
- **Description**: Delete a task by ID

### /project

- **Parameters For POST and PUT**:
    - `project_name` (string): Name of the project
    - `description` (string): Description of the project
    - `budget` (number): Budget for the project

- **Method**: GET
- **Description**: Get all projects

- **Method**: POST
- **Description**: Insert a new project (authentication required)

- **Method**: PUT
- **Description**: Update a project by ID

- **Method**: DELETE
- **Description**: Delete a project by ID