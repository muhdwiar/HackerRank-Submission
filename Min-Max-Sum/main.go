package main

import "fmt"

func miniMaxSum(arr []int32) {
	var value []int

	for i := 0; i < len(arr); i++ {
		var temp_sum int = 0
		for j := 0; j < len(arr); j++ {
			if j != i {
				temp_sum += int(arr[j])
			}

		}
		value = append(value, temp_sum)
	}

	min_val, max_val := value[0], value[0]
	for _, val := range value {
		if val > max_val {
			max_val = val
		}
		if val < min_val {
			min_val = val
		}
	}

	fmt.Println(min_val, max_val)

}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{7, 69, 2, 221, 8974})
}
