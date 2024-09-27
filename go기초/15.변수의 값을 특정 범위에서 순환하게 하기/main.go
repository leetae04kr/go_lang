package main

import "fmt"

func main() {
	const max_elem int = 5
	var i int = 0
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
}
