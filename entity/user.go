package entity

import (
	"time"
)

// use to display a user
type UserDisplay struct {
	ID				string		`json:"id"`
	Email			string		`json:"email"`
	Username		string		`json:"username" binding:"required"`
	FirstName		string		`json:"first_name" binding:"required"`
	LastName		string		`json:"last_name" binding:"required"`
	IsActive		bool		`json:"is_active"`
	IsStaff			bool		`json:"is_staff"`
	IsSuperuser		bool		`json:"is_superuser"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
	LastAuthenticatedAt		time.Time	`json:"last_authentication_at"`
}

// use to create or update a user
type UserCreateUpdate struct {
	UserDisplay
	Password		string		`json:"password"`
}

// Serializer to change a password
type ChangePassword struct {
	OldPassword		string		`json:"old_password" binding:"required"`
	NewPassword		string		`json:"new_password" binding:"required"`
}

// Used by a user to login
type UserLogin struct {
	Identifier	string	`json:"identifier" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

// Func that will check non empty field on UserDisplay and update user
func ValidateUpdate(user *UserCreateUpdate,u *UserDisplay) *UserCreateUpdate {
	if user.Email == "" {
		user.Email = u.Email
	}

	user.ID = u.ID
	user.Username = u.Username
	user.IsActive = u.IsActive
	user.IsStaff = u.IsStaff
	user.IsSuperuser = u.IsSuperuser
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
	user.LastAuthenticatedAt = u.LastAuthenticatedAt

	return user
}

func UserDisplayFormater(user *UserCreateUpdate) (u *UserDisplay) {
	u = &UserDisplay{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		IsActive: user.IsActive,
		IsStaff: user.IsStaff,
		IsSuperuser: user.IsSuperuser,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LastAuthenticatedAt: user.LastAuthenticatedAt,
	}

	return
}
