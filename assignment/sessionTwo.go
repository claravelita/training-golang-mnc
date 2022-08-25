package assignment

import (
	"fmt"

	"github.com/claravelita/training-golang-mnc/assignment/data"
	"github.com/claravelita/training-golang-mnc/assignment/dtos"
)

// Closure + Pointer + Struct
func ClosureStructAssignment() {
	persons := data.PersonName()
	printFriends := func(friends []*dtos.Person) {
		for _, p := range friends {
			fmt.Println(p.Name)
		}
	}

	printFriends(persons)
}
