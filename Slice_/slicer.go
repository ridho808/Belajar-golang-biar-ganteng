package slice

import "fmt"

func Slicer() {
	// tiga data slice Pointer,length,capacity
	// Pointer petunjuk array pertama
	// panjang array
	// capacity jumlah length dari pointer ke max array
	var Asalan = [...]string{"pisang", "mangga", "januari", "apa aja lah", "lagi males", "Ngapain bang", "Ganteng", "HAHA", "ialah"}

	Slice := Asalan[3:5]

	fmt.Println(Slice)

}
