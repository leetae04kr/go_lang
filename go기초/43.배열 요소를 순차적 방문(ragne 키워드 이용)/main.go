package main

import "fmt"

func main() {
	arr := [5]int{12, 34, 23, 56, 34}

	for i, value := range arr {
		fmt.Println(i, ":", value)
	}
}
