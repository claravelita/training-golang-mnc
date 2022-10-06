package session4

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claravelita/training-golang-mnc/assignment/helper"
	"github.com/claravelita/training-golang-mnc/assignment/session3"
	"github.com/claravelita/training-golang-mnc/dtos"
)

// Simple HTTP Server for Get and Register
type UserController struct {
	userService session3.ServiceUserInterface
}

func NewUserController(services session3.ServiceUserInterface) UserController {
	return UserController{
		userService: services,
	}
}

func (u *UserController) UserRoute() {
	http.HandleFunc("/register", u.RegisterHandler)
	http.HandleFunc("/user", u.GetHandler)
	http.HandleFunc("/", greet)
}

func (u *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dtos.User
	var response dtos.StandardResponse
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response = helper.BadRequestResponse()
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	u.userService.Register(&user)
	response = helper.SuccessResponse(nil)
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func (u *UserController) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response dtos.StandardResponse
	users := u.userService.GetUser()
	response = helper.SuccessResponse(users)
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func SimpleHTTPGetRegisterAssigment() {
	var PORT = ":8080"
	var db []*dtos.User
	initServiceUser := session3.NewServiceUser(db)
	initUserController := NewUserController(initServiceUser)
	initUserController.UserRoute()
	fmt.Println("Server Running on", PORT)
	http.ListenAndServe(PORT, nil)
}
