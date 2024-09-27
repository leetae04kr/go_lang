package main

import "fmt"

func main() {
	var scores [3]int = [3]int{1, 2}
	fmt.Println("=== scores ===")
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	var scores2 = [3]int{1, 2}
	fmt.Println("=== scores2 ===")
	for i := 0; i < len(scores2); i++ {
		fmt.Println(scores2[i])
	}

	scores3 := [3]int{1, 2}
	fmt.Println("=== scores3 ===")
	for i := 0; i < len(scores3); i++ {
		fmt.Println(scores3[i])
	}

	scores4 := [...]int{1, 2}
	fmt.Println("=== scores4 ===")
	for i := 0; i < len(scores4); i++ {
		fmt.Println(scores4[i])
	}

}
