package exercise

import "fmt"

func HipHura(angka int) {
	for i := 1; i <= angka; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("Hip Hura")
		} else if i%3 == 0 {
			fmt.Println("Hip")
		} else if i%5 == 0 {
			fmt.Println("Hura")
		} else {
			fmt.Println(i)
		}
	}
}
