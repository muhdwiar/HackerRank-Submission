package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	var big_sum int64

	for _, val := range ar {
		big_sum += val
	}

	return big_sum

}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
