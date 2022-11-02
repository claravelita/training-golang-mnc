package dtos

// [SESSION TWO] First init : Add name on Person struct
type Person struct {
	Name           string
	Address        string
	Job            string
	TrainingReason string
}

type User struct {
	Name string `json:"name"`
}

type Session11StatusResponse struct {
	Status Session11WinterWindResponse `json:"status"`
}

type Session11WinterWindResponse struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
