package assignment

import "fmt"

// Closure + Pointer + Struct
type Person struct {
	Name string
}

func ClosureStructAssignment() {
	persons := []*Person{
		{
			Name: "Clara",
		},
		{
			Name: "Fiqri",
		},
		{
			Name: "Medy",
		},
		{
			Name: "Ivan",
		},
		{
			Name: "Rijal",
		},
		{
			Name: "Adit",
		},
		{
			Name: "Luthfi",
		},
		{
			Name: "Tantut",
		},
		{
			Name: "Ian",
		},
		{
			Name: "Kemal",
		},
	}

	printFriends := func(friends []*Person) {
		for _, p := range friends {
			fmt.Println(p.Name)
		}
	}

	printFriends(persons)
}
