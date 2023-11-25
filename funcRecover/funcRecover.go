package funcrecover

// recover adalah function yang dapat digunakan untuk menangkap panic
// dengan recover panic akan terhenti sehingga program akan tetepa berjalan

import "fmt"

// function yang berfungsi menghentikan program / saat terjadi error
// defer functio tetap di eksekusi

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi error: ", message)
}

func RunApp(error bool) {
	defer endApp()
	if error {
		panic("UPS Error! recover")
	}

}
