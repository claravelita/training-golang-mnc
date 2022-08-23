package assignment

import "fmt"

func FirstAssignment() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "genap")
		} else {
			fmt.Println(i, "ganjil")
		}
	}
}
