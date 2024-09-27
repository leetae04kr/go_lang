package main

import "fmt"

func main() {
	const (
		JAN = iota + 1
		FEB
		MAR
		APR
		MAY
		JUN
		JUL
		AUG
		SEP
		OCT
		NOV
		DEC
	) //월
	fmt.Println("JAN:", JAN)
	fmt.Println("JUL:", JUL)
	fmt.Println("DEC:", DEC)

	const (
		c0 = iota * 10
		c1 = iota * 10
		c2 = iota * 10
		c3 = iota * 10
		c4 = iota * 10
	) //10씩 증가하는 상수
	fmt.Println("c1:", c1)
	fmt.Println("c1:", c2)
	fmt.Println("c1:", c3)
	fmt.Println("c1:", c4)
}
