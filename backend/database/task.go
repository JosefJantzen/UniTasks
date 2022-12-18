package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"desc"`
	Due         time.Time  `json:"due"`
	Done        bool       `json:"done"`
	DoneAt      *time.Time `json:"doneAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	UserId      uuid.UUID  `json:"userId"`
}

func (t *Task) merge(s *Task) {
	if s.Name != "" {
		t.Name = s.Name
	}
	if s.Description != "" {
		t.Description = s.Description
	}
	if s.Due != (time.Time{}) {
		t.Due = s.Due
	}
}

func (s *DBService) GetTaskById(id uuid.UUID, uid uuid.UUID) (*Task, error) {
	res, err := s.db.Query(
		"SELECT * FROM tasks WHERE id=$1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return nil, err
	}

	defer res.Close()
	res.Next()

	var tId uuid.UUID
	var name string
	var desc string
	var due time.Time
	var done bool
	var doneAt sql.NullTime
	var createdAt time.Time
	var updatedAt time.Time
	var userId uuid.UUID

	if err := res.Scan(&tId, &name, &desc, &due, &done, &doneAt, &createdAt, &updatedAt, &userId); err != nil {
		return nil, err
	}
	var doneAtPointer *time.Time = nil
	if doneAt.Valid {
		doneAtPointer = &doneAt.Time
	}
	task := Task{Id: tId, Name: name, Description: desc, Done: done, DoneAt: doneAtPointer, Due: due, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
	return &task, nil
}

func (s *DBService) GetTasksByUser(id uuid.UUID) ([]Task, error) {
	res, err := s.db.Query("SELECT * FROM tasks WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []Task{}

	for res.Next() {
		var tId uuid.UUID
		var name string
		var desc string
		var due time.Time
		var done bool
		var createdAt time.Time
		var updatedAt time.Time
		var userId uuid.UUID

		if err := res.Scan(&tId, &name, &desc, &due, &done, &createdAt, &updatedAt, &userId); err != nil {
			return nil, err
		}
		task := Task{Id: tId, Name: name, Description: desc, Due: due, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *DBService) InsertTask(task Task) (uuid.UUID, error) {
	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO tasks (name, due, description, user_id) VALUES ($1, $2, $3, $4) RETURNING id",
				task.Name,
				task.Due,
				task.Description,
				task.UserId,
			).Scan(&id)

			return err
		})
	return id, err
}

func (s *DBService) UpdateTask(reqTask Task) error {
	task, err := s.GetTaskById(reqTask.Id, reqTask.UserId)
	if err != nil {
		return err
	}
	task.merge(&reqTask)

	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE tasks SET name = $1, due = $2, description = $3, updated_at=now() WHERE id = $4 AND user_id=$5",
				task.Name,
				task.Due,
				task.Description,
				task.Id,
				task.UserId,
			)
			return err
		})
}

func (s *DBService) UpdateTaskDone(task Task) error {
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE tasks SET done = $1, updated_at=now() WHERE id = $2 AND user_id=$3",
				task.Done,
				task.Id,
				task.UserId,
			)
			return err
		})
}

func (s *DBService) DeleteTask(id uuid.UUID, userId uuid.UUID) error {
	res, err := s.db.Query(
		"DELETE FROM tasks WHERE id = $1 AND user_id=$2",
		id,
		userId,
	)
	if err != nil {
		return err
	}

	defer res.Close()
	return nil
}
