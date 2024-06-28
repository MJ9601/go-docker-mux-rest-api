package main

import "database/sql"

type Store interface {
	CreateUser() error
	GetUserByID(id string) (*User, error)

	CreateTask(t *Task) (*Task, error)
	GetTaskById(id string) (*Task, error)

	CreateProject(project *Project) (*Project, error)
	GetProjectById(id string) (*Project, error)

}

type Repository struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser() error {
	return nil
}

func (r *Repository) CreateTask(t *Task) (*Task, error){
	rows, err := r.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)
	if err != nil {
	 return nil, err
	}
	id, err := rows.LastInsertId()

	if err != nil{
		return nil, err
	}

	t.ID = id

	return t, nil
}

func (r *Repository) GetTaskById(id string) (*Task, error){
	var task Task
	 err := r.db.QueryRow("SELECT id, name, status, projectID, createdAt, assignedToID FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Name, &task.Status, &task.ProjectID, &task.CreatedAt, &task.AssignedToID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repository) GetUserByID(id string) (*User, error){
	var user User
	err := r.db.QueryRow("SELECT id, email, firstName, lastName, createdAt, password FROM users WHERE id = ?", id).Scan(&user.ID,  &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

