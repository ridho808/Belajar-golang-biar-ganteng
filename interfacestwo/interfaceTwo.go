package interfacestwo

import "fmt"

func interfaceKosong() interface{} {
	return "Ups"
}
func two() any {
	return (1)
}

func Ups() {
	ups := interfaceKosong()
	Any := two()
	fmt.Println(ups, Any)
}
