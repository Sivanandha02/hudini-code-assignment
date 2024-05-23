package main
//Simple, given a string of words, return the length of the shortest word(s).

//String will never be empty and you do not need to account for different data types.
import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
  str:=strings.Split(strings.TrimSpace(s)," ")
  short:=str[0]
  
  for i:=0;i<len(str)-1;i++{
    if len(short)>len(str[i+1]){
      short=str[i+1]
    }
  }
  return len(short)
  // your code
}
func main(){
	fmt.Println("shortest word",FindShort("the sojan philip"))
}