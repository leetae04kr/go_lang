package main

import "fmt"

func main() {
	var name string
	fmt.Print("이름 입력:")
	fmt.Scanln(&name)
	switch name {
	case "abc":
		fmt.Println("휘리릭~")
	case "ABC":
		fmt.Println("달려 달려~")
	default:
		fmt.Println("모르겠어요.")
	}
}
