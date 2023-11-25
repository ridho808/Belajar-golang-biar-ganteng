package tipemap

import "fmt"

func tipemap() {
	// map adalah tipe data berisikan kumpulan data yang sama namun kita bisa menentukan jenis tipedata index yang akan digunakan
	// map kumpulan key-value key harus uniq tidak boleh sama
	person := map[string]string{
		"name": "ridho",
		"umur": "24",
	}
	fmt.Println(person)
}
