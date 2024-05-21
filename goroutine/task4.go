package main

import "fmt"

func printNum(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	cha := make(chan int)
	go printNum(cha)
	for i:= range cha{
		println(i)
	}
	num := <-cha
	fmt.Println(num)
}