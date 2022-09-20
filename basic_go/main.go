package main

func main() {

	//first.go
	//gofuncs.First()

	//type.go
	//gofuncs.Type()

	//gopointer.go
	//gofuncs.Pointer()

	// functions.go
	/*
		x := 10
		y := 20
		fmt.Println(gofuncs.Addnum(x, y))

		// 가변인자함수
		gofuncs.Ssay("This", "is", "a", "book")
		gofuncs.Ssay("Hi")
	*/

	// ifexample.go
	//fmt.Println(gofuncs.Ppow(3, 2, 10), gofuncs.Ppow(3, 3, 20))

	// forloop.go
	/*
		gofuncs.Loop()

		gofuncs.Loop1()

		gofuncs.Printhello(5)
	*/

	// string.go
	//gofuncs.Addstr()

	// array.go
	/*
		gofuncs.SliceExam()
		gofuncs.SliceAdd()
	*/

	// map.go
	/*
		gofuncs.Mapexam()
		gofuncs.Mapforloop()
	*/

	/*
		//익명 함수 - 다른 함수의 인자나 반환값으로 사용
		allsum := func(n ...int) int {
			s := 0
			for _, i := range n {
				s += i
			}
			return s
		}
		fmt.Println(allsum(1, 2, 3, 4, 5))

		// 익명함수 예시 - add할수 선언후 인자 전달
		add := func(i int, j int) int {
			return i + j
		}
		r1 := gofuncs.Ccalc(add, 10, 20) // add 함수 전달
		fmt.Println(r1)

		// 직접 첫번째 파라미터에 익명함수 정의
		r2 := gofuncs.Ccalc(func(x int, y int) int { return x - y }, 10, 20)
		fmt.Println(r2)
	*/
}
