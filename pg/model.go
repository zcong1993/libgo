package pg

import (
	"time"
)

// Model is model for pg with uuid as primary key
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Model struct {
	ID        string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
