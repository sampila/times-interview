package main

import (
	"fmt"
	"time"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

func possibleTimes(digits []int) string {
	// Your code here
	var helper func([]int, int)
	res := [][]int{}

	helper = func(digits []int, n int) {
		if n == 1 {
			tmp := make([]int, len(digits))
			copy(tmp, digits)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(digits, n-1)
				if n%2 == 1 {
					tmp := digits[i]
					digits[i] = digits[n-1]
					digits[n-1] = tmp
				} else {
					tmp := digits[0]
					digits[0] = digits[n-1]
					digits[n-1] = tmp
				}
			}
		}
	}
	helper(digits, len(digits))

	var validTimes []string
	for _, t := range res {
		// If time is valid, print it.
		validTime, err := time.Parse("15:04", fmt.Sprintf("%d%d:%d%d", t[0], t[1], t[2], t[3]))
		if err != nil {
			continue
		} else {
			validTimes = append(validTimes, validTime.Format("15:04"))
		}
	}
	validUniqueTimes := uniqueTimes(validTimes)

	return fmt.Sprintf("Input: %v\nValid times: %+q\nReturn: %d\n", digits, validUniqueTimes, len(validUniqueTimes))
}

func uniqueTimes(arr []string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for e := range arr {
		// Check if already the mapped
		// variable is set to true or not.
		if !occurred[arr[e]] {
			occurred[arr[e]] = true
			// Append to result slice.
			result = append(result, arr[e])
		}
	}
	return result
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}
