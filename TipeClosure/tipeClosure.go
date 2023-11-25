package tipeclosure

import "fmt"

func tipeclosure() {
	// Kemampuan sebuah function berinteraksi dengan data disekitarnya dalam scope yang sama
	counter := 0

	increment := func() {
		counter++
	}
	increment()
	fmt.Println(counter)
}
