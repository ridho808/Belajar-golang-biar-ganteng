package nill

import "fmt"

func Names(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: name,
		}
	}
}

func CheckName() {
	name := Names("ridho")
	fmt.Println(name)
	if name == nil {
		fmt.Println("data Kosong")
	} else {
		fmt.Println("data terdapat value,", name["ridho"])
	}
}
