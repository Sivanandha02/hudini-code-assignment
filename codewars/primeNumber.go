package main

import "fmt"

func IsPrime(n int) bool {

	flag := 0
	if n <=1{
		return false
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			flag += 1
		}
	}
	if flag == 1 {
		return true
	}
	return false
}
func main() {
	
	fmt.Println(IsPrime(7))
}