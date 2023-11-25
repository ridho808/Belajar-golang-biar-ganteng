package forexpress

import "fmt"

func LoopingFor() {

	for length := 0; length < 10; length++ {
		fmt.Println("for1", length)
	}
	fmt.Println("\n")
	number := 0
	for number < 10 {
		number++
		fmt.Println("for2", number)
	}
	// array
	var name = [...]string{"ridho", "handika", "nigger", "NLE CHOPPA", "Jahseh", "Jared", "Miler"}

	for length := 0; length < len(name); length++ {

		fmt.Println("name=", name[length])
	}

	for index, list := range name {
		fmt.Println("name", index, list)
	}
	for _, list := range name {
		fmt.Println("name", list)
	}
}
