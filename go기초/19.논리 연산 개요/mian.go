package main

import "fmt"

func main() {
	fmt.Println("true&&true  =", true && true)
	fmt.Println("true&&false =", true && false)
	fmt.Println("false&&true =", false && true)
	fmt.Println("false&&false=", false && false)

	fmt.Println("true||true  =", true || true)
	fmt.Println("true||false =", true || false)
	fmt.Println("false||true =", false || true)
	fmt.Println("false||false=", false || false)

	fmt.Println("!false =", !false)
	fmt.Println("!true =", !true)

	fmt.Println("(2 < 3) && (3 < 4)", (2 < 3) && (3 < 4))
}
