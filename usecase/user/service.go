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
	hashedPassword, err := utils.HashWord(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return s.repo.Create(u)
}