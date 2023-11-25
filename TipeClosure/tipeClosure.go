package tipeclosure

import "fmt"

func TestClosures() {
	// Kemampuan sebuah function berinteraksi dengan data disekitarnya dalam scope yang sama
	counter := 0

	increment := func() {
		counter++
	}
	increment()
	increment()
	increment()
	fmt.Println(counter)
}
