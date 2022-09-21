package user

import (
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/utils"
)

type userservice struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *userservice {
	return &userservice{
		repo: r,
	}
}

func (s *userservice) List() ([]*entity.UserDisplay, error) {
	return s.repo.List()
}

func (s *userservice) Create(u *entity.UserCreateUpdate) error {
	hashedPassword, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}

// Retrieve a user
func (s *userservice) GetByID(id string) (*entity.UserDisplay, string, error) {
	u, password, err := s.repo.GetByID(id)
	if err != nil {
		return &entity.UserDisplay{}, "",entity.ErrNotFound
	}
	return u, password, nil
}

func (s *userservice) UpdateUser(u *entity.UserCreateUpdate) error {
	return s.repo.UpdateInfo(u)
}

func (s *userservice) UpdatePassword(u *entity.UserCreateUpdate) error {
	hashed_password, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed_password

	return s.repo.UpdatePassword(u)
}

func (s *userservice) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	return s.repo.SearchUser(identifier)
}

func (s *userservice) Delete(id string) error {
	return s.repo.Delete(id)
}