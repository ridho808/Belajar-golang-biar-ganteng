package array

import "fmt"

func Arrey() {
	var Names [3]string

	Names[0] = "HELLO"
	Names[1] = "World"
	Names[2] = "APAKABAR ?"

	fmt.Println(Names)

	var Angka = [3]string{
		"satu", "dua", "tiga",
	}
	fmt.Println(Angka)

	var Nomor = [5]int{
		1, 2, 3, 4, 5,
	}
	var Nomor2 = [5]int32{
		1339, 21234, 31222, 4333, 5112,
	}
	var Asalan = [...]string{"pisang", "mangga", "januari", "apa aja lah", "lagi males", "Ngapain bang", "Ganteng"}

	fmt.Println(Nomor)
	fmt.Println(Nomor2)
	fmt.Println("Array Nomor 1=", len(Nomor))
	fmt.Println("Array Asalan =", Asalan)
}
