package typeassertion

import "fmt"

func Assertion() interface{} {
	return 231
}

func CheckData() {
	result := Assertion()
	// resString := result.(string)
	// fmt.Println(resString)
	// error panic
	// resInt := result.(int)
	// fmt.Println(resInt)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)
	}
}
