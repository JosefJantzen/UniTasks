package database

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Due         time.Time `json:"due"`
	ParentUser  uuid.UUID `json:"parent"`
}

func (s *DBService) GetTaskById(id uuid.UUID) *Task {
	res, err := s.db.Query("SELECT * FROM tasks WHERE id=$1", id)
	if err != nil {
		return nil
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var name string
	var desc string
	var due time.Time
	var parent uuid.UUID

	if err := res.Scan(&uid, &name, &desc, &due, &parent); err != nil {
		return nil
	}
	task := Task{Id: uid, Name: name, Description: desc, Due: due, ParentUser: parent}
	return &task
}

func (s *DBService) GetTasksByUser(id uuid.UUID) []Task {
	res, err := s.db.Query("SELECT * FROM tasks WHERE parentUser=$1", id)
	if err != nil {
		return nil
	}

	defer res.Close()

	tasks := []Task{}

	for res.Next() {
		var id uuid.UUID
		var name string
		var due time.Time
		var desc string
		var parent uuid.UUID

		if err := res.Scan(&id, &name, &desc, &due, &parent); err != nil {
			return nil
		}
		task := Task{Id: id, Name: name, Due: due, Description: desc, ParentUser: parent}
		tasks = append(tasks, task)
	}
	return tasks
}
