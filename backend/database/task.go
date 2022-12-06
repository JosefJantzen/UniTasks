package database

import (
	"fmt"
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
		fmt.Println("asdad ", err)
		return nil
	}
	user := Task{Id: uid, Name: name, Description: desc, Due: due, ParentUser: parent}
	return &user
}

func (s *DBService) GetTasksByUser(id uuid.UUID) []Task {
	fmt.Println("543", id.String())
	res, err := s.db.Query("SELECT * FROM tasks WHERE parentUser=$1", id)
	if err != nil {
		fmt.Println("1 ", err)
		return nil
	}

	defer res.Close()

	tasks := []Task{}

	for res.Next() {
		fmt.Print("o")
		var id uuid.UUID
		var name string
		var due time.Time
		var desc string
		var parent uuid.UUID

		if err := res.Scan(&id, &name, &desc, &due, &parent); err != nil {
			fmt.Println("2 ", err)
			return nil
		}
		task := Task{Id: id, Name: name, Due: due, Description: desc, ParentUser: parent}
		tasks = append(tasks, task)
	}
	fmt.Println("t ", tasks)
	return tasks
}
