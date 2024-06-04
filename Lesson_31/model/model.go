package model

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID
	Name string
	Category string
	Cost int
}