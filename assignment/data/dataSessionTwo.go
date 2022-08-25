package data

import "github.com/claravelita/training-golang-mnc/assignment/dtos"

func PersonName() []*dtos.Person {
	persons := []*dtos.Person{
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
	return persons
}
