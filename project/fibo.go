package main

import "fmt"

func fiboNacci() {
	fmt.Println("enter the limit: ")
	var limit int
	fmt.Scanln(&limit)
	a:=0
	b:=1
	fmt.Println("fibonacci series...")
	fmt.Println(a)
	fmt.Println(b)
	for i:=2;i<limit;i++{
		temp:=a+b
		a=b
		b=temp
		fmt.Println(b)
	}
}
func main()  {
	fiboNacci()
}