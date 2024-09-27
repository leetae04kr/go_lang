package main

import "fmt"

func main() {
	var s []int

	fmt.Printf("용량:%d 원소 개수:%d\n", cap(s), len(s))
	var i int
	for i = 0; i < 10; i++ {
		s = append(s, i+1)
		fmt.Printf("용량:%d 원소 개수:%d\n", cap(s), len(s))
	}
}
