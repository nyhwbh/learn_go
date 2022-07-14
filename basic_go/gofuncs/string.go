package gofuncs

import "fmt"

func Addstr() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

}
