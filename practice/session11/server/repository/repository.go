package repository

import "github.com/claravelita/training-golang-mnc/practice/session11/server/model"

type UserRepository interface {
	FindById(id string) *model.User
}
