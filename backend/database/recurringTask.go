package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type RecurringTask struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Interval    int       `json:"interval"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      uuid.UUID `json:"userId"`
}

func (s *DBService) GetRecurringTaskById(id uuid.UUID) *RecurringTask {
	res, err := s.db.Query("SELECT * FROM recurring_tasks WHERE id=$1", id)
	if err != nil {
		return nil
	}

	defer res.Close()
	res.Next()

	var tId uuid.UUID
	var name string
	var desc string
	var start time.Time
	var end time.Time
	var interval int
	var createdAt time.Time
	var updatedAt time.Time
	var userId uuid.UUID

	if err := res.Scan(&tId, &name, &desc, &start, &end, &interval, &createdAt, &updatedAt, &userId); err != nil {
		return nil
	}
	task := RecurringTask{Id: tId, Name: name, Description: desc, Start: start, End: end, Interval: interval, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
	return &task
}

func (s *DBService) GetRecurringTasksByUser(id uuid.UUID) []RecurringTask {
	res, err := s.db.Query("SELECT * FROM recurring_tasks WHERE user_id=$1", id)
	if err != nil {
		return nil
	}

	defer res.Close()

	tasks := []RecurringTask{}

	for res.Next() {
		var tId uuid.UUID
		var name string
		var desc string
		var start time.Time
		var end time.Time
		var interval int
		var createdAt time.Time
		var updatedAt time.Time
		var userId uuid.UUID

		if err := res.Scan(&tId, &name, &desc, &start, &end, &interval, &createdAt, &updatedAt, &userId); err != nil {
			return nil
		}
		task := RecurringTask{Id: tId, Name: name, Description: desc, Start: start, End: end, Interval: interval, CreatedAt: createdAt, UpdatedAt: updatedAt, UserId: userId}
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *DBService) InsertRecurringTask(task RecurringTask) uuid.UUID {

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

	if err != nil {
		return uuid.Nil
	}
	return id
}

func (s *DBService) UpdateRecurringTask(task RecurringTask) error {
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
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
	if err != nil {
		return err
	}
	return nil
}

func (s *DBService) DeleteRecurringTask(id uuid.UUID, userId uuid.UUID) error {
	res, err := s.db.Query(
		"DELETE FROM recurring_tasks WHERE id = $1 AND user_id=$2",
		id,
		userId,
	)
	if err != nil {
		return err
	}

	defer res.Close()
	return err
}
