package stack

type Stack struct {
	stack_int []int
	stack_string []string
}

// int형 자료구조
// 데이터 저장 - int 배열에 값 추가
func (s *Stack) Push_i(item int){
	s.stack_int = append(s.stack_int, item)
}

// 데이터 삭제 - 길이가 0이면 Error 반환
func (s *Stack) Pop_i() (int,error) {
	stackLen := len(s.stack_int)

	if stackLen == 0 {
		return -1, error.New("No Data")
	}

	item, items := s.stack_int[stackLen-1], s.stack_int[0:stackLen-1]
	s.stack_int = items
	return item, nil
}

// string형 자료구조
// 데이터 저장 - string 배열에 값 추가
func (s *Stack) Push_s(item string){
	s.stack_string = append(s.stack_string, item)
}

// 데이터 삭제 - 길이가 0이면 Error 반환
func (s *Stack) Pop_s() (string,error) {
	stackLen := len(s.stack_string)

	if stackLen == 0 {
		return -1, error.New("No Data")
	}

	item, items := s.stack_string[stackLen-1], s.stack_string[0:stackLen-1]
	s.stack_string = items
	return item, nil
}