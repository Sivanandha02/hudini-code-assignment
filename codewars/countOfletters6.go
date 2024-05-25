package main

import (
	"fmt"
	"strings"
)

func countOfLetter(str string) {
	str1:= strings.ToLower(str)
	alph_count:=0
	count := map[string]int{}
	s:=strings.Split(str1,"")
	for _, v := range s {
		count[v]++
	}
	for i, value := range count {
		fmt.Println(i, value)
		if count[i]>1{
			alph_count++
		}
		
	}
	fmt.Println("count of letters : ",alph_count)
	
}
func main(){
	countOfLetter("aAbcde11")
}