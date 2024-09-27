package main

import "fmt"

func main() {
	var i int
	fmt.Println("Test 1:", i)
	i = 3 //i에 3을 대입
	fmt.Println("Test 2:", i)
	i += 2 //i = i+2
	fmt.Println("Test 3:", i)
	i -= 7 //i = i-7
	fmt.Println("Test 4:", i)
}
