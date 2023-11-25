package funcpanic

import "fmt"

// function yang berfungsi menghentikan program / saat terjadi error
// defer functio tetap di eksekusi

func endApp() {
	fmt.Println("End App")
}

func RunApp(error bool) {
	defer endApp()
	if error {
		panic("UPS Error!")
	}
}
