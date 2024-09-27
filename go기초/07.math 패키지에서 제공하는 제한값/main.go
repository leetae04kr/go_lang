package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt8)    //int8의 최솟값
	fmt.Println(math.MaxInt8)    //int8의 최댓값
	fmt.Println(math.MinInt16)   //int16의 최솟값
	fmt.Println(math.MaxInt16)   //int16의 최댓값
	fmt.Println(math.MinInt32)   //int32의 최솟값
	fmt.Println(math.MaxInt32)   //int32의 최댓값
	fmt.Println(math.MinInt64)   //int64의 최솟값
	fmt.Println(math.MaxInt64)   //int64의 최댓값
	fmt.Println(math.MaxFloat32) //float32의 최댓값
	fmt.Println(math.MaxFloat64) //float64의 최댓값
}
