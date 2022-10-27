package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	var hasil int32

	for _, val := range ar {
		hasil += val
	}

	return hasil

}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
}
