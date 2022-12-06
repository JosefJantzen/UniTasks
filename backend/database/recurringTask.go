package database

import "github.com/google/uuid"

type RecurringTask struct {
	id          uuid.UUID
	name        string
	description string
	interval    string
	parentUser  uuid.UUID
}
