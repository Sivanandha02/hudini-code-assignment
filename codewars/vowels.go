package main

import (
	"fmt"
	//"strings"
)


func SortVowels(str1 string) string {
  if str1 == ""{
    return ""
  }
  word:=""
  for i:=0;i<len(str1);i++{
    if string(str1[i])== "a" || string(str1[i])== "e" || string(str1[i])== "i" || string(str1[i])== "o" || string(str1[i])== "u" || string(str1[i])== "A" || string(str1[i])== "E" || string(str1[i])== "I" || string(str1[i])== "O" || string(str1[i])== "U"{
      word += "\n|" + string(str1[i])
      
    }else{
      word += "\n" + string(str1[i]) + "|" 
      
    }
  }
 return word[1:]
	
}
func main(){
	var str string
	fmt.Println("enter the string: ")
	fmt.Scanln(&str)
	fmt.Println(SortVowels(str))
}