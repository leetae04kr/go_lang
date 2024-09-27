package main

import "fmt"

func main() {
	var sum = 0
	var i int
	i = 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("1부터 100까지 정수 합계:", sum)
}
