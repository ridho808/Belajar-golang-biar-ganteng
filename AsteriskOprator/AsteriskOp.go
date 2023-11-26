package asteriskoprator

import "fmt"

type Address struct {
	City, Province, Country string
}

func AsteriskOP() {
	Address1 := Address{"Jakarta", "DKI JAKARTA", "Indonesia"}
	Address2 := &Address1
	Address2.City = "Bandung"
	fmt.Println(Address1)
	fmt.Println(Address2)
	Address2 = &Address{"Bandar Lampung", "Lampung", "Indonesia"}
	fmt.Println(Address2)
	var Address3 Address = Address{"Jogja", "Jawa Tengah", "Indonesia"}
	var Address4 *Address = &Address3
	fmt.Println(Address3)
	Address4.City = "Semarang"
	fmt.Println(Address3)
	fmt.Println(Address4)
	*Address4 = Address{"Texas", "Amrik Nggatau", "USA"}
	fmt.Println(Address3)
	fmt.Println(Address4)

}
