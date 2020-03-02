package user

import "fmt"

func foo() error {
	i := 0
	if i == 0 {
		return checkErr()
	}
	return nil
}

func checkErr() error {
	fmt.Println("123")
	return nil
}
