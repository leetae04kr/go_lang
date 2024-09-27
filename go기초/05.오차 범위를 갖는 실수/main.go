// Example 오차 범위를 갖는 실수
package main

import "fmt"

const errbound = 1e-14

func main() {
	var f1 = 0.0
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	f1 += 0.1
	fmt.Println(f1)
	fmt.Println(f1 == 1.0)              //오차 범위를 갖고 있으므로 버그일 수 있음
	fmt.Println((f1 - 1.0) <= errbound) //오차 범위내에 같은 값인지 비교
}

//Go 언어의 실수는 IEEE 754 규약을 따르며 오차 범위는 10의 -14승이예요. 따라서 실수 값을 갖는 변수와 비교하고자 하는 값을 뺀 후에 차이가 10의 -14승 이내에 있다면 오차 범위내에서 같은 값으로 취급할 수 있어요.
