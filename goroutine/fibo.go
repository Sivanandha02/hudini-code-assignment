package main

import "fmt"

func fiboNacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)

}
func main() {
	cha := make(chan int)
	go fiboNacci(10, cha)
	for i := range cha {
		fmt.Println(i)
	}
}