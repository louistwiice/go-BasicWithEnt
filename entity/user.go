package entity

import (
	"time"
)

// use to create or update a user
type UserCreateUpdate struct {
	ID				string		`json:"id"`
	Email			string		`json:"email"`
	FirstName		string		`json:"first_name" binding:"required"`
	LastName		string		`json:"last_name" binding:"required"`
	Password		string		`json:"password"`
	IsActive		bool		`json:"is_active"`
	IsStaff			bool		`json:"is_staff"`
	IsSuperuser		bool		`json:"is_superuser"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

// use to display a user
type UserDisplay struct {
	ID				string		`json:"id"`
	Email			string		`json:"email"`
	FirstName		string		`json:"first_name"`
	LastName		string		`json:"last_name"`
	IsActive		bool		`json:"is_active"`
	IsStaff			bool		`json:"is_staff"`
	IsSuperuser		bool		`json:"is_superuser"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

// Used by a user to login
type UserLogin struct {
	Identifier	string	`json:"identifier"`
	Password	string	`json:"password"`
}
