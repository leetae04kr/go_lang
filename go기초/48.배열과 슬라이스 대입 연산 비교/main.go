package main

import "fmt"

func main() {
	//배열
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]int

	fmt.Print("before arr1 :")
	fmt.Println(arr1)
	fmt.Print("before arr2 :")
	fmt.Println(arr2)

	arr2 = arr1

	arr2[0] = 8
	fmt.Print("after arr1 :")
	fmt.Println(arr1)
	fmt.Print("after arr2 :")
	fmt.Println(arr2)

	//슬라이스
	var s1 []int = []int{1, 2, 3}
	var s2 []int

	fmt.Print("before s1 :")
	fmt.Println(s1)
	fmt.Print("before s2 :")
	fmt.Println(s2)

	s2 = s1

	s2[0] = 8
	fmt.Print("after s1 :")
	fmt.Println(s1)
	fmt.Print("after s2 :")
	fmt.Println(s2)

}
