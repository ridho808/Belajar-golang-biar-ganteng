package funcdefer

import "fmt"

func Logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer Logging()
	fmt.Println("Aplikasi berjalan")
}

func FunctionDefer() {
	// adalah function yang dapat kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	// akan selalu di eksekusi walaupu  terjadi error
	runApplication()
}
