package main

import "fmt"

func main() {
	var name string
	fmt.Print("이름:")
	fmt.Scanln(&name)
	var num int
	fmt.Print("번호:")
	fmt.Scanln(&num)
	var addr string
	fmt.Print("주소:")
	fmt.Scanln(&addr)
	fmt.Println("이름은 ", name, " 번호는 ", num)
	fmt.Println("주소는 ", addr)
}
