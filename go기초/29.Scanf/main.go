package main

import "fmt"

func main() {
	var n1, n2, n3, n4 int
	fmt.Println("IPv4 주소")
	fmt.Scanf("%d.%d.%d.%d", &n1, &n2, &n3, &n4)
	fmt.Printf("입력한 IPv4주소는 %d.%d.%d.%d\n", n1, n2, n3, n4)
}
