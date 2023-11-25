package interfaces

import "fmt"

// interface adalah tipedata abstract tidak memiliki implementasi langsung
// sebuah interface berisikan definisi-definisi method
// biasanya inteface digunakan sebagai kontrak

type HasName interface {
	getName() string
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func SayHello(hasname HasName) {
	fmt.Println("HELLO", hasname.getName())
}

type Animal struct {
	name string
}

func (animal Animal) getName() string {
	return animal.name
}

func Cetak() {
	orang := Person{
		Name: "Ridho Ganteng",
	}
	hewan := Animal{
		name: "Biawak",
	}
	SayHello(hewan)
	SayHello(orang)
}
