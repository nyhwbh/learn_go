package gofuncs

import "fmt"

func Pointer() {
	var num int = 1

	var numPtr *int = &num

	fmt.Println(numPtr)
	fmt.Println(*numPtr)
	fmt.Println(&num)
}
