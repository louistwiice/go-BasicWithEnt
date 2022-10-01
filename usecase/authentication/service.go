package service_authentication

import (
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/utils"
)

type authservice struct {
	repo domain.UserRepository
}

func NewAuthService(r domain.UserRepository) *authservice {
	return &authservice{
		repo: r,
	}
}

func (s *authservice) Create(u *entity.UserCreateUpdate) error {
	hashedPassword, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}

// Service to update authentication date
func (s *authservice) UpdateAuthenticationDate(u *entity.UserDisplay) error {
	return s.repo.UpdateAuthenticationDate(u)
}

// Retrieve a user
func (s *authservice) GetByID(id string) (*entity.UserDisplay, string, error) {
	u, password, err := s.repo.GetByID(id)
	if err != nil {
		return &entity.UserDisplay{}, "",entity.ErrNotFound
	}
	return u, password, nil
}

func (s *authservice) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	return s.repo.SearchUser(identifier)
}