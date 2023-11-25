package recrusivefuntion

// recrusive function adalah function yang memanggil dirinya sendiri

func Factorial(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func FactorialTwo(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * FactorialTwo(number-1)
	}
}
