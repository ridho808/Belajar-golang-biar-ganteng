package tempstruct

import "fmt"

// struct adalah sebuah template data yang digunakan menggabungkan nol atau lebih tipedata lainnya dalam satu kesatuan
// struct biasanya representasi data dalam program aplikasi yang kita buat
//data stuct disimpan dalam field
// sederhanya struct adalah kumpulan dari field

type Customer struct {
	name, address string
	age           int
}

func TestStruct() {
	var Ridho Customer
	Ridho.name = "Ridho ganteng"
	Ridho.address = "Amerika selatan"
	Ridho.age = 24
	Megalodon := Customer{
		name:    "megawatlodon",
		address: "Jimbabway",
		age:     34,
	}
	fmt.Println(Ridho)
	fmt.Println(Megalodon)
}

// struct method : seolah struct memiliki fucntion
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello My name is", customer.name, name)
}

func HelloW() {
	handika := Customer{name: "handika"}
	handika.sayHello("Okey")
}
