package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var sum, temp_sum int32

	idx := 0
	for idx < len(arr) {
		sum += arr[idx][idx]
		idx++
	}

	temp_idx := 0
	for idx > 0 {
		idx--
		temp_sum += arr[temp_idx][idx]
		temp_idx++
	}

	sum = sum - temp_sum
	if sum < 0 {
		sum = sum * -1
	}

	return sum

}

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}
