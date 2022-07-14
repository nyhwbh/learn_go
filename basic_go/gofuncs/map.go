package gofuncs

import "fmt"

func Mapexam() {
	/*
		map[key자료형]value자료형{
			데이터
		}
	*/
	m := make(map[int]string)
	//	추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"
	//	키에 대한 값 읽기
	str := m[134]
	fmt.Println(str)
	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	fmt.Println(noData)
	// 삭제
	delete(m, 777)
}

func Mapforloop() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}
	// for range 문을 사용하여 모든 맵 요소 출력
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
