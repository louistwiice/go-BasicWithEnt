package repository

import (
	"context"

	"github.com/google/uuid"
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
		Select(user.FieldID, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldIsActive, user.FieldIsStaff, user.FieldIsSuperuser, user.FieldCreatedAt, user.FieldUpdatedAt).
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

func (c *UserClient) GetByID(id string) (*entity.UserDisplay, string, error) {
	var u entity.UserDisplay
	ctx := context.Background()
	id_convert, err := uuid.Parse(id) // Convert the string to uuid type
	if err != nil {
		return nil, "", err
	}

	resp := c.client.User.
		Query().
		Where(user.ID(id_convert)).
		AllX(ctx)

	if len(resp) > 0 {
		u.ID = resp[0].ID.String()
		u.Email = resp[0].Email
		u.FirstName = resp[0].FirstName
		u.LastName = resp[0].LastName
		u.IsActive = resp[0].IsActive
		u.IsStaff = resp[0].IsStaff
		u.IsSuperuser = resp[0].IsSuperuser
		u.CreatedAt = resp[0].CreatedAt
		u.UpdatedAt = resp[0].UpdatedAt	
	} else {
		return nil, "", entity.ErrNotFound
	}
	

	return &u, resp[0].Password, nil
}

// Update user information, except password
func (c *UserClient) UpdateInfo(u *entity.UserCreateUpdate) error {
	ctx := context.Background()
	id_convert, err := uuid.Parse(u.ID) // Convert the string to uuid type
	if err != nil {
		return err
	}

	_, err = c.client.User.UpdateOneID(id_convert).
			SetFirstName(u.FirstName).
			SetLastName(u.LastName).
			SetEmail(u.Email).
			Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Update user password
func (c *UserClient) UpdatePassword(u *entity.UserCreateUpdate) error {
	ctx := context.Background()
	id_convert, err := uuid.Parse(u.ID) // Convert the string to uuid type
	if err != nil {
		return err
	}

	_, err = c.client.User.UpdateOneID(id_convert).
			SetPassword(u.Password).
			Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

//Search a user information by email or username
func (c *UserClient) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	var u entity.UserDisplay
	ctx := context.Background()

	resp := c.client.User.
		Query().
		Where(
			user.Or(user.Email(identifier)),
		).
		AllX(ctx)

	if len(resp) > 0 {
		u.ID = resp[0].ID.String()
		u.Email = resp[0].Email
		u.FirstName = resp[0].FirstName
		u.LastName = resp[0].LastName
		u.IsActive = resp[0].IsActive
		u.IsStaff = resp[0].IsStaff
		u.IsSuperuser = resp[0].IsSuperuser
		u.CreatedAt = resp[0].CreatedAt
		u.UpdatedAt = resp[0].UpdatedAt	
	} else {
		return nil, "", entity.ErrNotFound
	}
	
	return &u, resp[0].Password, nil
}
