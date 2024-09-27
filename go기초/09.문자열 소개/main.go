package main

import "fmt"

func main() {
	var s1 string = "Hello, google.co.kr" //이중 콤마로 표현
	var s2 string = "안녕, 언제나 휴일"

	//백쿼터(~자판에 있는 ')로 여러 줄 표현
	var s3 = `안녕
    언제나 휴일`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
