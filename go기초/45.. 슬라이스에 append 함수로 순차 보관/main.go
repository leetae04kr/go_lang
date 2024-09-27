// 5명 성적은 슬라이스에 초기 설정, 5명의 성적은 입력받기
package main

import "fmt"

func main() {
	var arr [5]int = [5]int{90, 88, 76, 80, 99}
	var scores []int = make([]int, 5, 10)
	var i int

	for i = 0; i < 5; i++ {
		scores[i] = arr[i]
	}

	fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	var score int
	for i = 5; i < 10; i++ {
		fmt.Printf("%d 번 성적:", i+1)
		fmt.Scanln(&score)
		scores = append(scores, score)
		fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	}

	fmt.Println("==== 학생 성적 출력 ====")

	var sum int
	for i = 0; i < 10; i++ {
		sum += scores[i]
		fmt.Printf("%d번 성적:%d\n", i+1, sum)
	}
	fmt.Printf("총점:%d\n", sum)
}
