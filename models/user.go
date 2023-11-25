package models

import (
	"time"
)

type User struct {
	ID        int        `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Surname   string     `db:"surname" json:"surname"`
	Birthdate time.Time  `db:"birthdate" json:"birthdate"`
	Username  string     `db:"username" json:"username"`
	Password  string     `db:"password" json:"password"`
	CreatedAt time.Time  `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
	CreatedBy *int       `db:"createdBy" json:"createdBy"`
	UpdatedBy *int       `db:"updatedBy" json:"updatedBy"`
}

type UserReq struct {
	Name      string `db:"name" json:"name"`
	Surname   string `db:"surname" json:"surname,omitempty"`
	Birthdate string `db:"birthdate" json:"birthdate"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password,omitempty"`
}
