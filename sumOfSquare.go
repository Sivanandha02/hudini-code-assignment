package main
func SquareSum(numbers []int) int {
	sum:=0
	for i:=0;i<len(numbers);i++{
	  square:=numbers[i]*numbers[i]
	  sum+=square
	}
	return sum
	  // your code here
  }