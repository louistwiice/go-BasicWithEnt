package user

import (
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/utils"
)

type service struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) List() ([]*entity.UserDisplay, error) {
	return s.repo.List()
}

func (s *service) Create(u *entity.UserCreateUpdate) error {
	hashedPassword, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}

// Retrieve a user
func (s *service) Get(id string) (*entity.UserDisplay, string, error) {
	u, password, err := s.repo.Get(id)
	if err != nil {
		return &entity.UserDisplay{}, "",entity.ErrNotFound
	}
	return u, password, nil
}

func (s *service) UpdateUser(u *entity.UserCreateUpdate) error {
	return s.repo.UpdateInfo(u)
}

func (s *service) UpdatePassword(u *entity.UserCreateUpdate) error {
	hashed_password, err := utils.HashString(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed_password

	return s.repo.UpdatePassword(u)
}