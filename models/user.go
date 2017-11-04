package models

import "time"

type User struct {
	ID             int64     `json:"id" form:"id"`
	Username       string    `json:"username" form:"username"`
	HashedPassword []byte    `json:"-" form:"-"`
	CreatedAt      time.Time `json:"created_time" form:"created_time"`
}