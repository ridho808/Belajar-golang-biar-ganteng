package string

import "fmt"

func Str() {
	var Nama string = "Ridho"
	var Face string = "Ganteng"
	var GameFav string = "Mobile Ganggu"

	fmt.Println("Nama =", Nama)
	fmt.Println("face =", Face)
	fmt.Println("Game Favorite =", GameFav)

	// function for string

	var LengthStr int = len(Nama)
	fmt.Println("Panjang String =", LengthStr)
	var GetI = Nama[1]
	fmt.Println(string(GetI))
}
