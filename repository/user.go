package repository

import (
	"context"

	"github.com/louistwiice/go/basicwithent/ent"
	"github.com/louistwiice/go/basicwithent/ent/user"
	"github.com/louistwiice/go/basicwithent/entity"
)


type UserClient struct {
	client *ent.Client
}

func NewUserClient(client *ent.Client) * UserClient {
	return &UserClient{
		client: client,
	}
} 

// List all users
func (c *UserClient) List() ([]*entity.UserDisplay, error) {
	var u []*entity.UserDisplay
	ctx := context.Background()

	err := c.client.User.
		Query().
		Select(user.FieldID, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldIsActive, user.FieldIsStaff, user.FieldIsSuperuser, user.FieldCreatedAt, user.FieldCreatedAt).
		Scan(ctx, &u)

	if err != nil {
		return nil, err
	}
	
	return u, nil
}

// Create a user
func (c *UserClient) Create(u *entity.UserCreateUpdate) error {
	ctx := context.Background()

	resp, err := c.client.User.
		Create().
		SetEmail(u.Email).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetPassword(u.Password).
		SetIsActive(u.IsActive).
		SetIsStaff(u.IsStaff).
		SetIsSuperuser(u.IsSuperuser).
		Save(ctx)

	if err != nil {
		return err
	}

	u.ID = resp.ID.String()
	u.CreatedAt = resp.CreatedAt
	u.UpdatedAt = resp.UpdatedAt
	return nil
}