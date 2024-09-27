package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("두 개의 정수 입력:")
	fmt.Scanln(&x, &y)

	switch {
	case x > y:
		fmt.Println("큰 값:", x)
	case x < y:
		fmt.Println("큰 값:", y)
	default:
		fmt.Println("두 수는 서로 같습니다.")
	}
}
