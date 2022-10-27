package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var score []int32

	count := 0
	for count < 2 {
		score = append(score, 0)
		count++
	}

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				score[0] += 1
			} else if a[i] < b[i] {
				score[1] += 1
			}
		}

	}

	return score
}

func main() {
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
}
