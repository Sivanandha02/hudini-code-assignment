package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()
	go func() {
		ch <- "World!"
	}()
	msg1:= <- ch
	msg2:= <- ch
	fmt.Println(msg1)
	fmt.Println(msg2)
}
