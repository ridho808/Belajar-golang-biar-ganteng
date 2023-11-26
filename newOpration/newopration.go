package newopration

import "fmt"

type Address struct {
	city, province, country string
}

func CheckAddress() {
	address1 := new(Address)
	address2 := address1

	address2.country = "Indonesia"
	fmt.Println(address1, address2)
}
