package assignment

import (
	"fmt"

	"github.com/claravelita/training-golang-mnc/assignment/dtos"
)

// Assignment bikin interface Get and Create
type userService struct {
	db []*dtos.User
}

type ServiceUserInterface interface {
	Register(request *dtos.User) string
	GetUser() []*dtos.User
}

func NewServiceUser(db []*dtos.User) ServiceUserInterface {
	return &userService{db: db}
}

func (u *userService) Register(request *dtos.User) string {
	u.db = append(u.db, request)
	return fmt.Sprint(request.Name, " berhasil didaftarkan")
}

func (u *userService) GetUser() []*dtos.User {
	return u.db
}
