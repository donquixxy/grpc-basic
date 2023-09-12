package repo

import (
	"gokit-basic/services/user/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(v *domain.UserDomain) (*domain.UserDomain, error)
	GetListUser() []*domain.UserDomain
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (s *userRepo) CreateUser(v *domain.UserDomain) (*domain.UserDomain, error) {

	res := s.db.Create(v)

	return v, res.Error
}

func (s *userRepo) GetListUser() []*domain.UserDomain {
	var v []*domain.UserDomain

	s.db.Find(&v)

	return v
}
