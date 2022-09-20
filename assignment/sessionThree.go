package assignment

import (
	"fmt"
	"sync"

	"github.com/claravelita/training-golang-mnc/data"
	"github.com/claravelita/training-golang-mnc/dtos"
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

func InterfaceGetCreateAssigment() {
	var user []*dtos.User
	initUserService := NewServiceUser(user)
	listPerson := data.PersonName()

	for _, p := range listPerson {
		res := initUserService.Register(&dtos.User{Name: p.Name})
		fmt.Println(res)

	}

	resGet := initUserService.GetUser()
	fmt.Println("\n-----------Hasil Get User-------------")
	for _, v := range resGet {
		fmt.Println(v.Name)
	}
}

func InterfaceGetCreateWithGoRoutineAssigment() {
	var user []*dtos.User
	initUserService := NewServiceUser(user)
	listPerson := data.PersonName()

	var wg sync.WaitGroup
	for _, p := range listPerson {
		wg.Add(1)

		go func(name string) {
			res := initUserService.Register(&dtos.User{Name: name})
			defer wg.Done()
			fmt.Println(res)
		}(p.Name)

	}
	wg.Wait()

	resGet := initUserService.GetUser()
	fmt.Println("\n-----------Hasil Get User-------------")
	for _, v := range resGet {
		fmt.Println(v.Name)
	}
}
