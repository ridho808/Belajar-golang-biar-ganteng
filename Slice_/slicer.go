package slice

import "fmt"

func Slicer() {
	// tiga data slice Pointer,length,capacity
	// Pointer petunjuk array pertama
	// panjang array
	// capacity jumlah length dari pointer ke max array
	var Asalan = [...]string{"pisang", "mangga", "januari", "apa aja lah", "lagi males", "Ngapain bang", "Ganteng", "HAHA", "ialah"}
	var names = [6]string{"ridho", "handika", "kundul", "rhoma", "akbar", "zorin"}
	var slice1 []string = names[3:]
	Slice := Asalan[3:5]
	tambah := append(slice1, "mantap jiwa cok")
	fmt.Println(Slice)
	fmt.Println(slice1)
	fmt.Println(tambah)
}
