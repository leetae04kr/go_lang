package main

import "fmt"

func main() {
	var name, addr string
	var num int
	fmt.Println("이름 번호 주소")
	var re int
	re, _ = fmt.Scanln(&name, &num, &addr)
	fmt.Println("이름은 ", name, " 번호는 ", num)
	fmt.Println("주소는 ", addr)
	fmt.Println(re)
}
