package main

import (
	integer "belajar-golang/Integer"
	Slice "belajar-golang/Slice_"
	Str "belajar-golang/String"
	array "belajar-golang/arr"
	Bool "belajar-golang/boolean"
	"belajar-golang/forexpress"
	helloWlrd "belajar-golang/helloworld"
	ifexpression "belajar-golang/ifExpression"
	recrusivefuntion "belajar-golang/recrusiveFuntion"
	switchexpres "belajar-golang/switchExpres"
	tipedeklarasi "belajar-golang/tipedeklarasi"
	"belajar-golang/tipemap"
	"fmt"
)

// function : block kode yg dibuat untuk dieksekusi secara berulang //

func printSting() {
	fmt.Println("Ini string dari function printSting")
}

// function parameter
func Tambah(angka1 int, angka2 int) {
	var a int = angka1
	var b int = angka2
	var c int = a + b
	fmt.Println("print variale C=", c)
}

// function return value : mengembalikan data(function harus mendeklarasikan tipedata yang akan di hasilkan)
func Greeting(name string) string {
	return "hello " + name
}

// fuction multiple value : mengembalikan multi value (wajib menyebutkan semua return valuenya)
func getFullname() (string, string) {
	return "Ridho", "Ganteng"
}

func getLastName() (string, string) {
	return "Ganteng banget", "Ridho"
}

// named function multivalues : memberikan variable di return value
func getcompliteName() (firstname, middlename, lastname string) {
	firstname = "Ridho"
	middlename = "Syam"
	lastname = "ganteng banget"
	return firstname, middlename, lastname
}

// variadic functiun : varargs bisa menerima lebih dari satu input atau semacam array (paling ujung parameter)
// parameter array harus membuat arraynya terlebih dahulu
// parameter menggunakan varargs bisa langsung mengirim datanya,jika lebih darisatu maka menggunakan koma
func SumAll(number ...int) int {
	total := 0
	for _, angka := range number {
		total += angka
	}
	return total
}

type Filter func(string) string

// function as parameter : function sebagai parameter di function
func sayHelloWithFilter(name string, filter Filter) {
	filtername := filter(name)
	fmt.Println("Hello", filtername)
}

func spamFilter(name string) string {
	if name == "babi" || name == "anjing" {
		return "...."
	} else {
		return name
	}
}

// anonim funciton / function tanpa nama
type BlackList func(string) bool

func Register(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	helloWlrd.Hello()
	integer.Int_()
	Bool.Boolean()
	Str.Str()
	array.Arrey()
	Slice.Slicer()
	tipedeklarasi.TypeDeclar()
	tipemap.TypeMap()
	ifexpression.IfExp()
	switchexpres.SwitchExp()
	forexpress.LoopingFor()
	printSting()
	Tambah(1, 3)
	// ++++++++++++++++++++++++++++ //
	fmt.Println(Greeting("ridho"))

	firstName, LastName := getFullname()

	fmt.Println(firstName, LastName)

	_, first := getLastName()
	fmt.Println(first)
	a, b, c := getcompliteName()
	fmt.Println(a, b, c)

	var sum int = SumAll(10, 20, 30, 40, 50, 60)
	fmt.Println(sum)
	// function as parameters
	sayHelloWithFilter("Ridho", spamFilter)
	sayHelloWithFilter("babi", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
	// anon func
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	Register("Ridho ganteng", blacklist)
	Register("Anjing", func(name string) bool {
		return name == "babi" || name == "Anjing"
	})
	recrusive := recrusivefuntion.Factorial(10)
	fmt.Println(recrusive)
	recrusivetwo := recrusivefuntion.FactorialTwo(10)
	fmt.Println(recrusivetwo)
}
