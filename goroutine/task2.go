// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
package main

import (
	//"sync"
	"time"
)
type waitGrp struct{
	add int
	done int
}
func (wg *waitGrp) Add(){
	wg.add++
	
}
func (wg *waitGrp) Done(){
	wg.add--
}
func (wg *waitGrp) Wait()  {
	for {
		if wg.add==0{
			break
		}
	}
}
func launchNum(wg *waitGrp){
	for i:=1;i<=5;i++{
		println(i)
		time.Sleep(time.Second)
	}
	defer wg.Done()
 }

 func main(){
	var w waitGrp
   
 
    for i := 0; i < 3; i++ {
        w.Add()
        go launchNum(&w)
 
    }
    w.Wait()
 
}
	// w:=waitGrp{}
	// w.Add(1)

	// go launchNum(w)
	// go launchNum()
	// w.Done()
	// // creating waitGroup
	//var wg sync.WaitGroup

	// call three go routines
	 //for i:=0;i<3;i++{
	// 	wg.Add(1)
	// 	go launchNum(&wg)
	// }

	//wait for all three goroutines
	//wg.Wait()
 