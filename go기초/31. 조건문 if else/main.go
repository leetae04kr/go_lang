package main

import "fmt"

func main() {
	var num int
	fmt.Print("번호:")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println(num, "은(는) 짝수")
	} else {
		fmt.Println(num, "은(는) 홀수")
	}
	fmt.Println("테스트 종료")
}
