package postgres

import (
	"fmt"

	"github.com/claravelita/training-golang-mnc/practice/session11/server/model"
	"github.com/claravelita/training-golang-mnc/practice/session11/server/repository"
)

type userRepo struct{}

func NewUserRepo() repository.UserRepository {
	return &userRepo{}
}

func (u *userRepo) FindById(id string) *model.User {
	query := "SELECT * FROM users WHERE id=$1"

	fmt.Println(query)
	return nil
}
