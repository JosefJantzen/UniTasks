package database

import (
	"fmt"

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
