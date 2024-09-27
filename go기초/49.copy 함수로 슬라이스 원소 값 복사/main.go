package main

import "fmt"

func main() {
	var s1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 []int = make([]int, 2, 5)
	s2[0] = 8
	s2[1] = 9
	fmt.Printf("s1 용량 %d 원소 개수:%d\n", cap(s1), len(s1))
	fmt.Println(s1)
	fmt.Printf("s2 용량 %d 원소 개수:%d\n", cap(s2), len(s2))
	fmt.Println(s2)

	copy(s2, s1)

	fmt.Println("===copy(s2,s1) 수행 후===")
	fmt.Printf("s1 용량 %d 원소 개수:%d\n", cap(s1), len(s1))
	fmt.Println(s1)
	fmt.Printf("s2 용량 %d 원소 개수:%d\n", cap(s2), len(s2))
	fmt.Println(s2)

}
