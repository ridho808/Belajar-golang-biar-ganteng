package pointerfunc

import "fmt"

type Address struct {
	city, province, country string
}

func changeCountry(address *Address) {
	address.country = "Indoneisa"
	fmt.Println(address)
}

func PointerFunc() {
	address := &Address{"Padang", "sumatra barat", ""}
	changeCountry(address)
	fmt.Println(address)
}
