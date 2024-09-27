package main

import "fmt"

func main() {
	arr := [5]int{12, 34, 23, 56, 34}
	var length int
	length = len(arr)
	for i := 0; i < length; i++ {
		fmt.Println(i, ":", arr[i])
	}
}
