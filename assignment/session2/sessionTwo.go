package session2

import (
	"fmt"
	"os"
	"strconv"

	"github.com/claravelita/training-golang-mnc/data"
	"github.com/claravelita/training-golang-mnc/dtos"
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

func CommandLineArgumentAssignment() {
	param := os.Args[1]
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Error! parameter integer only")
		return
	}
	persons := data.PersonName()
	countPersons := len(persons)
	if paramInt > countPersons || paramInt < 1 {
		fmt.Println("Error! range parameter only", 1, "-", countPersons)
		return
	}

	paramInt = paramInt - 1

	fmt.Println("- Name:", persons[paramInt].Name)
	fmt.Println("- Address:", persons[paramInt].Address)
	fmt.Println("- Job:", persons[paramInt].Job)
	fmt.Println("- Training Reason:", persons[paramInt].TrainingReason)
}
