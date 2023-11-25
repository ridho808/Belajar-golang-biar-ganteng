package ifexpression

import "fmt"

func IfExp() {
	name := "Ridho"
	if name == "Ridho" {
		fmt.Println("hai Name =", name)
	} else {
		fmt.Println("Name !=", false)
	}
	name = "Akbars"
	if name == "Ridho" {
		fmt.Println("Haii Ridho")
	} else if name == "Akbar" {
		fmt.Println("Hallo Akbar")
	} else if name == "Muhammad" {
		fmt.Println("Hallo Muhammad")
	} else {
		fmt.Println("Duh dak kenal aku deck")
	}
	// sortStatment
	if length := len(name); length > 5 {
		fmt.Println("name panjang lebih dari 5 huruf", length)
	} else if length <= 5 {
		fmt.Printf("name kurang dari 5 huruf")
	}
}
