package main

func main(){
	//stack 사용방법
	//int형
	s:= Stack{stack_int}
	s.Push(1)
	s.Push(2)
	item,err := s.Pop()
	if err != nil{
		fmt.Println(err)
	}
}