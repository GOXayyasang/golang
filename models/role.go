package models

import "time"

type Role struct {
	ID        int        `json:"id" db:"id" mapstructure:"id"`
	Name      string     `json:"name" db:"name" mapstructure:"name"`
	Status    bool       `json:"status" db:"status" mapstructure:"status"`
	CreatedBy int        `json:"created_by" db:"created_by" mapstructure:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" mapstructure:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at" mapstructure:"updated_at"`
	UpdatedBy *int       `json:"updated_by" db:"updated_by" mapstructure:"updated_by"`
	Token     string     `json:"token" mapstructure:"token"`
}
