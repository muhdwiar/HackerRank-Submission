package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	var count, max_val int32

	for _, val1 := range candles {
		if val1 > max_val {
			max_val = val1
		}
	}

	for _, val2 := range candles {
		if val2 == max_val {
			count += 1
		}
	}

	return count
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}
