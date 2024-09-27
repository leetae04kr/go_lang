// Example 복소수 표현
package main

import "fmt"

func main() {
	var c1 complex64 = 1 + 2i //실수부와 허수부
	var c2 complex64 = 3i     //허수부만 있을 때
	var c3 complex64 = 2.0    //실수부만 있을 때
	fmt.Println(c1 + c2)      //복소수 덧셈
	fmt.Println(c1 - c3)      //복소수 뺄셈

	var f1 = real(c1) //실수부 추출
	var f2 = imag(c1) //허수부 추출
	fmt.Println(f1)
	fmt.Println(f2)

	var c4 = complex(f1, f2) //complex 함수로 복소수 합성
	fmt.Println(c4)
}
