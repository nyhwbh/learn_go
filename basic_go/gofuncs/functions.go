package gofuncs

import "fmt"

func Addnum(x int, y int) int {
	return x + y
}

//가변인자 함수 - 입력되는 인자 값이 몇개인지 알 수 없을 때 사용
func Ssay(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

//익명함수 선언시 함수의 정보(입력값 수, 출력값 수 등을 입력해 줘야 한다.)
func Ccalc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
