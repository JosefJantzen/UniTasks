package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
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

func (s *DBService) InsertTask(task Task) uuid.UUID {

	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO tasks (name, due, description, parentUser) VALUES ($1, $2, $3, $4) RETURNING id",
				task.Name,
				task.Due,
				task.Description,
				task.ParentUser,
			).Scan(&id)

			if err != nil {
				return err
			}
			return nil
		})

	if err != nil {
		return uuid.Nil
	}

	return id
}
