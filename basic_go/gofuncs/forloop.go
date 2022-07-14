package gofuncs

import "fmt"

func Loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Loop1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Printhello(x int) int {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Hello World")
		}
	}
	return x
}
