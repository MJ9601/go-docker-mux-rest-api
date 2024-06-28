package main

type MockStore struct {
}

func (m *MockStore) CreateUser() error {
	return nil
}

func (m *MockStore) CreateTask(task *Task) (*Task, error) {
	return nil, nil
}

func (m *MockStore) GetTaskById(id string) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetUserByID(id string) (*User, error) {
	return &User{}, nil
}