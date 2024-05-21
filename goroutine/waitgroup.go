package main

import "sync"

type waitGrp1 struct{
	add int
	done int
}
func (wg *waitGrp) Add(){
	wg.add+=1
	
}
func (wg *waitGrp) Done(){
	wg.done--
}
func main() {
	a:=waitGrp()
}