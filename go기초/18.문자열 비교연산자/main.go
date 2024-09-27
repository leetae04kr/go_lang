package main

import "fmt"

func main() {
	var s1, s2, s3 string = "abc", "zb", "ABC"
	fmt.Println("a의 아스키 코드:", 'a')
	fmt.Println("z의 아스키 코드:", 'z')
	fmt.Println("A의 아스키 코드:", 'A')

	fmt.Println(s1, "<", s2, " 결과", s1 < s2)
	fmt.Println(s1, "<", s3, " 결과", s1 < s3)

	var hs1, hs2, hs3 string = "강감찬", "킹콩", "홍길동"
	fmt.Println(hs1, "<", hs2, " 결과", hs1 < hs2)
	fmt.Println(hs1, "<", hs3, " 결과", hs1 < hs3)
}
