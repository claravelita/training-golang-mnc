package assignment

import "fmt"

// 1 - Assignment Condition and Looping
func ConditionLoopingAssignment() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "genap")
		} else {
			fmt.Println(i, "ganjil")
		}
	}
}

// 2 - Slice + Looping + Kenalan
func SliceLoopingAssignment() {
	listName := []string{"Clara", "Fiqri", "Medy", "Ivan", "Rijal", "Adit", "Luthfi", "Tantut", "Ian", "Kemal"}
	for _, i := range listName {
		fmt.Println(i)
	}
}
