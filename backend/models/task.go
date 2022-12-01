package models

import "github.com/google/uuid"

type Task struct {
	id          uuid.UUID
	name        string
	description string
	due         string
	parentUser  uuid.UUID
}
