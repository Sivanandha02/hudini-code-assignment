package main
func GetSum(a, b int) int {
	sum:=0
	min:=a
	max:=b
	if a>b{
	  min=b
	  max=a
	}
	for i:=min;i<=max;i++{
	  sum+=i
	
	}
	return sum
  }