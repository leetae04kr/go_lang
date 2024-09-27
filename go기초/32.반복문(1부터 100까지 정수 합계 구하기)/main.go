package main

import "fmt"

func main() {
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1부터 100까지 정수 합계:", sum)
}
