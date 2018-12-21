package ginhelper

import "time"

// ErrResp is error response
type ErrResp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// Model is same as gorm.Model but with json tag
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
