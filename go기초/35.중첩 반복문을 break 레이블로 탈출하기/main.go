package main

import "fmt"

func main() {
	var sum = 0
	var num int
Loop1:
	for i := 0; i < 10; i++ {
	Loop2:
		for j := 0; j < 10; j++ {
			fmt.Print("정수 입력:")
			fmt.Scanln(&num)
			if num < 0 {
				break Loop1
			} else {
				if num > 100 {
					break Loop2
				}
			}
			sum += num
			fmt.Println("j:", j)
		}
		fmt.Println("i:", i)
	}
	fmt.Println("합계:", sum)

}
