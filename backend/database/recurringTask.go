package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type RecurringTask struct {
	Id          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"desc"`
	Start       time.Time              `json:"start"`
	Ending      time.Time              `json:"ending"`
	Interval    int                    `json:"interval"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	UserId      uuid.UUID              `json:"userId"`
	History     []RecurringTaskHistory `json:"history"`
}

func (t *RecurringTask) merge(s *RecurringTask) {
	if s.Name != "" {
		t.Name = s.Name
	}
	if s.Description != "" {
		t.Description = s.Description
	}
	if s.Start != (time.Time{}) {
		t.Start = s.Start
	}
	if s.Ending != (time.Time{}) {
		t.Start = s.Start
	}
	if s.Interval != 0 {
		t.Interval = s.Interval
	}
}

func (s *DBService) CheckRecurringTaskExists(id uuid.UUID, uid uuid.UUID) bool {
	res, err := s.db.Query(
		"SELECT id FROM recurring_tasks WHERE id=$1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return false
	}
	defer res.Close()
	res.Next()
	var Id uuid.UUID

	if err := res.Scan(&Id); err != nil {
		return false
	}
	return true
}

func (s *DBService) GetRecurringTaskById(id uuid.UUID, uid uuid.UUID) (*RecurringTask, error) {
	res, err := s.db.Query(
		"SELECT * FROM recurring_tasks WHERE id=$1 AND user_id=$2",
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
	var start time.Time
	var ending time.Time
	var interval int
	var createdAt time.Time
	var updatedAt time.Time
	var userId uuid.UUID

	if err := res.Scan(&tId, &name, &desc, &start, &ending, &interval, &createdAt, &updatedAt, &userId); err != nil {
		return nil, err
	}
	task := RecurringTask{Id: tId, Name: name, Description: desc, Start: start, Ending: ending, Interval: interval, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
	task.History, err = s.GetRecurringTasksHistory(task.Id, task.UserId)
	return &task, err
}

func (s *DBService) GetRecurringTasksByUser(id uuid.UUID) ([]RecurringTask, error) {
	res, err := s.db.Query("SELECT * FROM recurring_tasks WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []RecurringTask{}

	for res.Next() {
		var tId uuid.UUID
		var name string
		var desc string
		var start time.Time
		var ending time.Time
		var interval int
		var createdAt time.Time
		var updatedAt time.Time
		var userId uuid.UUID

		if err := res.Scan(&tId, &name, &desc, &start, &ending, &interval, &createdAt, &updatedAt, &userId); err != nil {
			return nil, err
		}
		task := RecurringTask{Id: tId, Name: name, Description: desc, Start: start, Ending: ending, Interval: interval, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
		task.History, err = s.GetRecurringTasksHistory(task.Id, task.UserId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *DBService) InsertRecurringTask(task RecurringTask) (uuid.UUID, error) {
	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO recurring_tasks (name, interval, description, user_id) VALUES ($1, $2, $3, $4) RETURNING id",
				task.Name,
				task.Interval,
				task.Description,
				task.UserId,
			).Scan(&id)

			if err != nil {
				return err
			}
			return nil
		})
	return id, err
}

func (s *DBService) UpdateRecurringTask(reqTask RecurringTask) error {
	task, err := s.GetRecurringTaskById(reqTask.Id, reqTask.UserId)
	if err != nil {
		return err
	}
	task.merge(&reqTask)
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE recurring_tasks SET name = $1, interval = $2, description = $3, updated_at=now() WHERE id = $4 AND user_id=$5",
				task.Name,
				task.Interval,
				task.Description,
				task.Id,
				task.UserId,
			)
			return err
		})
}

func (s *DBService) DeleteRecurringTask(id uuid.UUID, userId uuid.UUID) error {
	task, err := s.GetRecurringTaskById(id, userId)
	if err != nil {
		return err
	}
	if task.Id != id {
		return errors.New("unauthorized deletion")
	}
	s.DeleteCompleteRecurringTaskHistory(id, userId)
	res, err := s.db.Query(
		"DELETE FROM recurring_tasks WHERE id = $1 AND user_id=$2",
		id,
		userId,
	)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}
