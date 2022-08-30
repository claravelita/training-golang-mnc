package main

import (
	"fmt"

	"github.com/claravelita/training-golang-mnc/assignment"
	"github.com/claravelita/training-golang-mnc/assignment/data"
	"github.com/claravelita/training-golang-mnc/assignment/dtos"
)

func main() {
	// assignment.ConditionLoopingAssignment()
	// assignment.SliceLoopingAssignment()
	// assignment.ClosureStructAssignment()
	// assignment.CommandLineArgumentAssignment()

	var user []*dtos.User
	initUserService := assignment.NewServiceUser(user)
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
