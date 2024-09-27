package main

import "fmt"

func main() {
	var b1 bool = true //값으로 true 혹은 false

	fmt.Println(b1)

	b1 = 2 > 3 //비교 연산의 결과
	fmt.Println(b1)

	b1 = true && false //논리 연산의 결과
	fmt.Println(b1)
}
