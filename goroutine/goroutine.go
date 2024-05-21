package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _,value:= s{
		sum ++
	}
	c <- sum

}
func main() {
	s := []int{1, 2, 3, 4, 2, 1}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	x := <-c
	fmt.Println(x)
}