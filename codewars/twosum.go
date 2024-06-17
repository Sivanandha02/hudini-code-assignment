package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

func main() {
	numbers := []int{2, 3, 4, 7, 8}
	target := 9
	fmt.Println(TwoSum(numbers, target))
}