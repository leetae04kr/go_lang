package main

import "fmt"

func main() {
	var score int
	fmt.Print("점수 입력:")
	fmt.Scanln(&score)
	if (score < 0) || (score > 100) {
		fmt.Println("잘못 입력하셨네요.")
		return
	}

	switch score / 10 {
	case 10:
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("B")
	case 6:
		fmt.Println("B")
	case 5:
		fmt.Println("B")

	default:
		fmt.Println("모르겠어요.")
	}
}
