package data

import "github.com/claravelita/training-golang-mnc/assignment/dtos"

func PersonName() []*dtos.Person {
	persons := []*dtos.Person{
		{
			Name:           "Clara",
			Address:        "Bekasi",
			Job:            "Backend Developer",
			TrainingReason: "Ingin lebih paham best practice pada golang",
		},
		{
			Name:           "Fiqri",
			Address:        "Bandung",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Medy",
			Address:        "Jakarta",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Ivan",
			Address:        "Tangerang",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Rijal",
			Address:        "Pinrang",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Adit",
			Address:        "Jakarta",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Luthfi",
			Address:        "Tegal",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Tantut",
			Address:        "Lampung",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Ian",
			Address:        "Lampung",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
		{
			Name:           "Kemal",
			Address:        "Lampung",
			Job:            "Backend Developer",
			TrainingReason: "Training",
		},
	}
	return persons
}
