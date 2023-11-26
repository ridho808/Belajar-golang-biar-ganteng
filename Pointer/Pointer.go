package pointer

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer menggunakan & untuk pass by value

func Array1() {
	address := Address{"Lampung", "bandar Lampung", "Indonesia"}
	address2 := &address
	address2.City = "Tangerang"

	var alamat Address = Address{"Tangerang", "Banten", "Indonesia"}
	var alamat2 *Address = &alamat
	alamat2.City = "Bekasi"

	fmt.Println(address)
	fmt.Println(address2)
	fmt.Println(alamat)
	fmt.Println(alamat2)
}
