// package declaration
package main

//  import packages
import (
	"fmt"
	"math"
	"math/rand"
)

// function declaration
func hellow() {
	// print
	fmt.Println("Hello World!")
}

/* Random numbers */
func random() {
	fmt.Println("random no is ", rand.Intn(10))
}

/* sqrt */
func square() {
	fmt.Printf("now %g problems.\n", math.Sqrt(7))
}

// add wo numbers

func main() {
	hellow()
	random()
	square()
	fmt.Println("sum is ", sum(40, 10))

	a, b := StringPrint("cat", "dog")
	fmt.Println(a, b)

	fmt.Println(split(10))
	fmt.Println("Aveage of array ", calculateAverage())
	fmt.Println(checkAge(25))
	fmt.Println(reverseString("hello"))
	fmt.Printf("%d is largest number\n", largestNumber(numbers))
//
//initialize count in COunter struct value 
// if we dont initilize it default value is 0
	c:=Counter{}
	c1:=add(add(c))
	fmt.Println(c1)
	c2:=decrement(c1)
	fmt.Println(c2)	
	c3:=reset(c2)
	fmt.Println(c3)

// monkey problem
fmt.Println(monkeyCount(10))
// make negative problem
fmt.Println(MakeNegative(-1))
// count
fmt.Println(CountBy(3,10))
//power of two
fmt.Println(PowersOfTwo(3))
// POints
fmt.Println(Points([]string{"3:1","1:3","3:3"}))
//getsum
fmt.Println(GetSum(-1,2))
//shortest string
fmt.Println(FindShort("good morning su"))
//sum of squares of numbers
fmt.Println("sum of square of numbers ",SquareSum([]int{3,4,5,1}))
//count of odd numbers
fmt.Println("count of odd numbers",OddCount(15))
}
