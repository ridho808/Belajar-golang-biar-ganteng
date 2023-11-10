package main

import (
	integer "belajar-golang/Integer"
	Slice "belajar-golang/Slice_"
	Str "belajar-golang/String"
	array "belajar-golang/arr"
	Bool "belajar-golang/boolean"
	"belajar-golang/exercise"
	helloWlrd "belajar-golang/helloworld"
	tipedeklarasi "belajar-golang/tipedeklarasi"
	"fmt"
)

func main() {
	helloWlrd.Hello()
	integer.Int_()
	Bool.Boolean()
	Str.Str()
	array.Arrey()
	Slice.Slicer()
	tipedeklarasi.TypeDeclar()
	// exercise
	fmt.Println("_+++++++++++_________++++++++")
	exercise.HipHura(15)
}
