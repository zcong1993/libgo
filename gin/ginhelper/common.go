package ginhelper

import "time"

// ErrResp is error response
type ErrResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
