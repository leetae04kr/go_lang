package main

import "fmt"

func main() {
	var n int
	fmt.Print("학생수:")
	fmt.Scanln(&n)

	var scores []int = make([]int, n)
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%d 번 성적:", i+1)
		fmt.Scanln(&scores[i])
	}

	fmt.Println("==== 학생 성적 출력 ====")

	var sum int
	for i = 0; i < n; i++ {
		sum += scores[i]
		fmt.Printf("%d번 성적:%d\n", i+1, sum)
	}
	fmt.Printf("총점:%d\n", sum)

}
