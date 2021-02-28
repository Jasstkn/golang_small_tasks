package main

import (
	"fmt"
	"math"
)

const baseNum = 2

func main() {
	for _, v := range pawInt(10) {
		fmt.Println(v)
	}
}

func pawInt(n int) []float64 {
	var resultArray []float64

	for i := 1; i <= n; i++ {
		resultArray = append(resultArray, math.Pow(baseNum, float64(i)))
	}

	return resultArray
}
