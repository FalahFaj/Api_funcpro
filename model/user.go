package model

import "time"

type User struct {
	Id        int64  `json:"id" db:"id"`
	Nama      string `json:"nama" db:"nama"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Role      string `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}