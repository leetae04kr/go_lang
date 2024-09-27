package main

import "fmt"

func main() {
	var scores [3]int
	fmt.Println("=== scores ===")
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}
