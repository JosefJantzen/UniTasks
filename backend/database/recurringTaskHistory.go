package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type RecurringTaskHistory struct {
	Id              uuid.UUID  `json:"id"`
	Description     string     `json:"desc"`
	Due             time.Time  `json:"due"`
	Done            bool       `json:"done"`
	DoneAt          *time.Time `json:"doneAt"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UserId          uuid.UUID  `json:"userId"`
	RecurringTaskId uuid.UUID  `json:"recurringTaskId"`
}

func (t *RecurringTaskHistory) merge(s *RecurringTaskHistory) {
	if s.Description != "" {
		t.Description = s.Description
	}
}

func (s *DBService) GetRecurringTaskHistoryById(id uuid.UUID, uid uuid.UUID) (*RecurringTaskHistory, error) {
	res, err := s.db.Query(
		"SELECT * FROM recurring_tasks_history WHERE id=$1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return nil, err
	}

	defer res.Close()
	res.Next()

	var tId uuid.UUID
	var desc string
	var due time.Time
	var done bool
	var doneAt sql.NullTime
	var createdAt time.Time
	var updatedAt time.Time
	var userId uuid.UUID
	var recurringTaskId uuid.UUID

	if err := res.Scan(&tId, &desc, &due, &done, &doneAt, &createdAt, &updatedAt, &userId, &recurringTaskId); err != nil {
		return nil, err
	}
	var doneAtPointer *time.Time = nil
	if doneAt.Valid {
		doneAtPointer = &doneAt.Time
	}
	task := RecurringTaskHistory{Id: tId, Description: desc, Due: due, Done: done, DoneAt: doneAtPointer, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId, RecurringTaskId: recurringTaskId}
	return &task, nil
}

func (s *DBService) GetRecurringTaskHistoriesByUser(id uuid.UUID) ([]RecurringTaskHistory, error) {
	res, err := s.db.Query("SELECT * FROM recurring_tasks_history WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []RecurringTaskHistory{}

	for res.Next() {
		var tId uuid.UUID
		var desc string
		var due time.Time
		var done bool
		var doneAt sql.NullTime
		var createdAt time.Time
		var updatedAt time.Time
		var userId uuid.UUID
		var recurringTaskId uuid.UUID

		if err := res.Scan(&tId, &desc, &due, &done, &doneAt, &createdAt, &updatedAt, &userId, &recurringTaskId); err != nil {
			return nil, err
		}
		var doneAtPointer *time.Time = nil
		if doneAt.Valid {
			doneAtPointer = &doneAt.Time
		}
		task := RecurringTaskHistory{Id: tId, Description: desc, Due: due, Done: done, DoneAt: doneAtPointer, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId, RecurringTaskId: recurringTaskId}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *DBService) GetRecurringTasksHistory(id uuid.UUID, uid uuid.UUID) ([]RecurringTaskHistory, error) {
	res, err := s.db.Query(
		"SELECT * FROM recurring_tasks_history WHERE recurring_task_id=$1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []RecurringTaskHistory{}

	for res.Next() {
		var tId uuid.UUID
		var desc string
		var due time.Time
		var done bool
		var doneAt sql.NullTime
		var createdAt time.Time
		var updatedAt time.Time
		var userId uuid.UUID
		var recurringTaskId uuid.UUID

		if err := res.Scan(&tId, &desc, &due, &done, &doneAt, &createdAt, &updatedAt, &userId, &recurringTaskId); err != nil {
			return nil, err
		}
		var doneAtPointer *time.Time = nil
		if doneAt.Valid {
			doneAtPointer = &doneAt.Time
		}
		task := RecurringTaskHistory{Id: tId, Description: desc, Due: due, Done: done, DoneAt: doneAtPointer, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId, RecurringTaskId: recurringTaskId}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *DBService) InsertRecurringTaskHistory(task RecurringTaskHistory) (uuid.UUID, error) {
	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO recurring_tasks_history (description, due, done, doneAt, user_id, recurring_task_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
				task.Description,
				task.Due,
				task.Done,
				task.DoneAt,
				task.UserId,
				task.RecurringTaskId,
			).Scan(&id)

			if err != nil {
				return err
			}
			return nil
		})
	return id, err
}

func (s *DBService) UpdateRecurringTaskHistory(reqTask RecurringTaskHistory) error {
	task, err := s.GetRecurringTaskHistoryById(reqTask.Id, reqTask.UserId)
	if err != nil {
		return err
	}
	task.merge(&reqTask)
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE recurring_tasks_history SET description = $1, due = $2, updated_at=now() WHERE id = $3 AND recurring_task_id=$4 AND user_id=$5",
				task.Description,
				task.Due,
				task.Id,
				task.RecurringTaskId,
				task.UserId,
			)
			return err
		})
}

func (s *DBService) UpdateRecurringTaskHistoryDone(task RecurringTaskHistory) error {
	return crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE recurring_tasks_history SET done = $1, done_at=$2, updated_at=now() WHERE id = $3 AND recurring_task_id=$4 AND user_id=$5",
				task.Done,
				task.DoneAt,
				task.Id,
				task.RecurringTaskId,
				task.UserId,
			)
			return err
		})
}

func (s *DBService) DeleteRecurringTaskHistory(id uuid.UUID, uid uuid.UUID) error {
	res, err := s.db.Query(
		"DELETE FROM recurring_tasks_history WHERE id = $1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}

func (s *DBService) DeleteCompleteRecurringTaskHistory(id uuid.UUID, uid uuid.UUID) error {
	res, err := s.db.Query(
		"DELETE FROM recurring_tasks_history WHERE recurring_task_id=$1 AND user_id=$2",
		id,
		uid,
	)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}
