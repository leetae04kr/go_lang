package main

import "fmt"

func main() {
	var origin_s []int = []int{10, 23, 34, 47, 54, 62, 7, 89, 91, 102}
	var start int
	var last int

	fmt.Print("원본 슬라이스:")
	fmt.Println(origin_s)

	fmt.Println("원본 슬라이스에서 부분 슬라이스 만들기")
	fmt.Print("시작 인덱스:")
	fmt.Scanln(&start)
	fmt.Print("끝 인덱스:")
	fmt.Scanln(&last)

	var sub_s []int = origin_s[start : last+1]
	fmt.Print("부분 슬라이스:")
	fmt.Println(sub_s)

}
