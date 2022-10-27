package main

import (
	"fmt"
)

func gradingStudents(grades []int32) []int32 {

	for idx, curr_grade := range grades {
		if curr_grade >= 38 {
			if (curr_grade+2)%5 == 0 {
				curr_grade += 2
			}

			if (curr_grade+1)%5 == 0 {
				curr_grade += 1
			}
			grades[idx] = curr_grade
		}

	}
	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
}
