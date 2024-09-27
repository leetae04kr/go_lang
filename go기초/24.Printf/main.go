package main

import "fmt"

func main() {
	fmt.Printf("2<3=%b\n", 2 < 3)          //bool
	fmt.Printf("23은 이진수로 %b\n", 23)        //이진수
	fmt.Printf("family name %c\n", '홍')    //문자
	fmt.Printf("10진수 출력:%d\n", 23)         //10진수
	fmt.Printf("8진수 출력:%o\n", 23)          //8진수
	fmt.Printf("16진수 출력:%x\n", 23)         //16진수,a~f
	fmt.Printf("16진수 출력:%X\n", 23)         //16진수,A~F
	fmt.Printf("고정소수점:%f\n", 123.4567)     //고정소수점
	fmt.Printf("고정소수점:%F\n", 123.4567)     //고정소수점
	fmt.Printf("지수 표현:%e\n", 123.4567)     //지수 표현, e
	fmt.Printf("지수 표현:%E\n", 123.4567)     //지수 표현, E
	fmt.Printf("간단한 실수 표현:%g\n", 123.4567) //간단한 실수 표현
	fmt.Printf("간단한 실수 표현:%g\n", 123.4567) //간단한 실수 표현
	fmt.Printf("문자열:%s", "홍길동\n")          //문자열
	var a int = 32
	fmt.Printf("메모리주소:%p\n", a)                     //포인터
	fmt.Printf("유니코드 %U\n", '\ud55c')               //유니코드
	fmt.Printf("%T\n", 23)                          //타입
	fmt.Printf("모든 형식:%v , %v\n", 23, 'a')          //모든 형식
	fmt.Printf("형식도 함께:%d %#o , %#x\n", 23, 23, 23) //%#v 형식을 구분할 수 있게
	fmt.Printf("%4d %04d\n", 23, 23)                //출력 폭 지정, 빈 곳 0출력
	fmt.Printf("%-4d %-4d\n", 23, 23)               //왼쪽 정렬
	fmt.Printf("%9.2f\n", 123.4567)                 //소수점 이하 자리 출력 지
}
