package gofuncs

import "fmt"

func Arrayexam() {
	//[...]으로 배열을 선언하면 {}안에 있는 원소 수 만큼의 크기를 가지는 배열이 생성된다.
	//한번 생성된 배열의 크기는 변경할 수 없다.
	fruits := [...]string{"사과", "바나나", "포도"}
	fmt.Println(fruits[1])

	fruits[1] = "수박"
	fmt.Println(fruits[1])
}

func SliceExam() {
	//slice 선언방법
	//길이가 10인 슬라이스 생성
	sliceFruits := make([]string, 10)

	//길이가 5인 슬라이스
	sliceFruits2 := []string{"사과", "바나나", "포도", "수박", "망고"}

	sliceFruits[0] = "사과"
	sliceFruits[1] = "바나나"
	sliceFruits[2] = "토마토"

	fmt.Println(sliceFruits)
	fmt.Println(sliceFruits2)

	//슬라이스 자르기 [시작점:끝점]
	fmt.Println(sliceFruits2[1:3])
}

func SliceAdd() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...) // 이어붙이기
	// 토마토를 제외하고 이어붙이기
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

}
