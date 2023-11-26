package pointermethod

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "mr ." + man.Name
}
func PointerMethod() {
	ridho := Man{"ridho"}
	ridho.Married()
	fmt.Println(ridho)
}
