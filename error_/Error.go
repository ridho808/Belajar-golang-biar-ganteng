package error_

import (
	"errors"
	"fmt"
)

func testError(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi return Nol ")
	} else {
		return nilai / pembagi, nil
	}
}

func Pembagian() {
	result, err := testError(12, 0)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(result)
	}
}
