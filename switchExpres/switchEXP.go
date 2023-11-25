package switchexpres

import "fmt"

func SwitchExp() {
	name := "Gak tau"
	switch name {
	case "Ridho":
		fmt.Println("Nama Ridho")
	case "Eko":
		fmt.Println("Nama Eko")
	case "Babang":
		fmt.Println("Nama Babang")
	default:
		fmt.Println("Gak Tau ah")
	}
}
