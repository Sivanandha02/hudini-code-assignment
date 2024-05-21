package main
func MakeNegative(x int) int {
	switch{
	  case x>=1:
		return -x
	  case x<0:
		return x
	  default:
		return 0
	}
  }