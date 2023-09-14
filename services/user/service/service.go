package service

import (
	"errors"
	"gokit-basic/services/user/domain"
	"gokit-basic/services/user/repo"
)

type UserService interface {
	CreateUser(v *domain.UserDomain) (*domain.UserDomain, error)
	GetListUser() ([]*domain.UserDomain, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(v *domain.UserDomain) (*domain.UserDomain, error) {
	return s.repo.CreateUser(v)
}

func (s *userService) GetListUser() ([]*domain.UserDomain, error) {
	data := s.repo.GetListUser()

	if len(data) == 0 {
		return nil, errors.New("user not found")
	}

	return data, nil
}
