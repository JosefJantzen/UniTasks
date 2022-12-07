package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
)

type RecurringTask struct {
	Id          uuid.UUID
	Name        string
	Description string
	Interval    int
	ParentUser  uuid.UUID
}

func (s *DBService) GetRecurringTaskById(id uuid.UUID) *RecurringTask {
	res, err := s.db.Query("SELECT * FROM recurring_tasks WHERE id=$1", id)
	if err != nil {
		fmt.Println("11", err)
		return nil
	}

	defer res.Close()
	res.Next()

	var uid uuid.UUID
	var name string
	var desc string
	var interval int
	var parent uuid.UUID

	if err := res.Scan(&uid, &name, &desc, &interval, &parent); err != nil {
		fmt.Println("22", err)
		return nil
	}
	task := RecurringTask{Id: uid, Name: name, Description: desc, Interval: interval, ParentUser: parent}
	return &task
}

func (s *DBService) GetRecurringTasksByUser(id uuid.UUID) []RecurringTask {
	res, err := s.db.Query("SELECT * FROM recurring_tasks WHERE parentUser=$1", id)
	if err != nil {
		fmt.Println("2", err)
		return nil
	}

	defer res.Close()

	tasks := []RecurringTask{}

	for res.Next() {
		var uid uuid.UUID
		var name string
		var desc string
		var interval int
		var parent uuid.UUID

		if err := res.Scan(&uid, &name, &desc, &interval, &parent); err != nil {
			fmt.Println("1", err)
			return nil
		}
		task := RecurringTask{Id: uid, Name: name, Description: desc, Interval: interval, ParentUser: parent}
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *DBService) InsertRecurringTask(task RecurringTask) uuid.UUID {

	var id uuid.UUID
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			err := tx.QueryRow(
				"INSERT INTO recurring_tasks (name, interval, description, parentUser) VALUES ($1, $2, $3, $4) RETURNING id",
				task.Name,
				task.Interval,
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

func (s *DBService) UpdateRecurringTask(task RecurringTask) error {
	err := crdb.ExecuteTx(context.Background(), s.db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"UPDATE recurring_tasks SET name = $1, interval = $2, description = $3 WHERE id = $4",
				task.Name,
				task.Interval,
				task.Description,
				task.Id,
			)
			return err
		})
	if err != nil {
		return err
	}
	return nil
}
