package main

import "fmt"

func main() {
	var i1 int = 2
	var i2 int = 3

	fmt.Println(i1, "+", i2, "=", i1+i2)
	fmt.Println(i1, "-", i2, "=", i1-i2)
	fmt.Println(i1, "*", i2, "=", i1*i2)

	var f1 float32 = 0.2
	var f2 float32 = 0.3
	fmt.Println(f1, "+", f2, "=", f1+f2)
	fmt.Println(f1, "-", f2, "=", f1-f2)
	fmt.Println(f1, "*", f2, "=", f1*f2)

	var c1 complex64 = 2 + 3i
	var c2 complex64 = 1 + 2i
	fmt.Println(c1, "+", c2, "=", c1+c2)
	fmt.Println(c1, "-", c2, "=", c1-c2)
	fmt.Println(c1, "*", c2, "=", c1*c2)

	var s1 string = "hello,"
	var s2 string = " ehclub.co.kr"
	fmt.Println(s1, "+", s2, "=", s1+s2)
}
