//Write a method (or function, depending on the language) that converts a string to camelCase,
// that is, all words must have their first letter capitalized and spaces must be removed.

//Examples (input --> output):
// "hello case" --> "HelloCase"
// "camel case word" --> "CamelCaseWord"
package main

import (
	"fmt"
	"strings"
	//"unicode"
)
func CamelCase(s string) string {
    words := strings.Fields(s)
    for i := range words{
      words[i]=strings.Title(words[i])
    }
  camel:= strings.Join(words,"")
  return camel
}

func main(){
  st:="hello world"
  fmt.Println("original : ",st)
  fmt.Println("camelcase : ",CamelCase(st))
}