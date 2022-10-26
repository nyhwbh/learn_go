// 질문
// 트리보나치 함수 만들기
// 트리보나치는 피보나치함수의 연장선으로 앞에 있는 3개의 정수를 합한 값을 출력한다.
// 2번째 array의  n의 갯수만큼 출력한다.
// 단, 정수는 0 이상의 양수만 해당한다.
// 예시
// tribonacci([0, 0, 1], 9) // 결과값 == [0, 0, 1, 1, 2, 4, 7, 13, 24]
// tribonacci([1, 1, 1], 5) // 결과값 == [1, 1, 1, 3, 5]
// tribonacci([1, 2, 3], 2) // 결과값 == [1, 2]

package eZar

import "fmt"

func Print() {
	item := []int{0, 0, 1}
	item2 := []int{1, 1, 1}
	item3 := []int{1, 2, 3}
	fmt.Println(tribonacci(item, 9))
	fmt.Println(tribonacci(item2, 5))
	fmt.Println(tribonacci(item3, 2))

}

func tribonacci(item []int, num int) []int {
	answer := make([]int, num)
	if num < 3 {
		for i := 0; i < num; i++ {
			answer[i] = item[i]
		}
		return answer
	} else {
		answer[0] = item[0]
		answer[1] = item[1]
		answer[2] = item[2]
		for i := 3; i < num; i++ {
			count := answer[i-3] + answer[i-2] + answer[i-1]
			answer[i] = count
		}
	}

	return answer
}
