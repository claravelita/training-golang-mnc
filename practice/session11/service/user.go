package service

import (
	"database/sql"

	"github.com/claravelita/training-golang-mnc/practice/session11/server/model"
	"github.com/claravelita/training-golang-mnc/practice/session11/server/repository"
)

type UserServices struct {
	repo repository.UserRepository
}

func NewUserServices(repo repository.UserRepository) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) GetSingleUserById(id string) (*model.User, error) {
	user := u.repo.FindById(id)
	if user == nil {
		return nil, sql.ErrNoRows
	}

	product := u.repo.FindById(id)
	if product == nil {
		return nil, sql.ErrNoRows
	}

	return user, nil
}
