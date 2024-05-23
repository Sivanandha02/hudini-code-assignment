package main

import "fmt"

//Given two integers a and b, which can be positive or negative,
// find the sum of all the integers between and including them and return it.
//If the two numbers are equal return a or b.
func GetSum(a, b int) int {
	sum:=0
	min:=a
	max:=b
	if a==b{
	  return a
	}
	if a>b{
	  min=b
	  max=a
	}else{
	  min=a
	  max=b
	}
	for i:=min;i<=max;i++{
	  sum+=i
	}
	return sum
  }
  func main(){
	fmt.Println("sum: ",GetSum(10,20))
  }