package pg

import (
	"github.com/google/uuid"
	"time"
)

// Model is model for pg with uuid as primary key
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Model struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
