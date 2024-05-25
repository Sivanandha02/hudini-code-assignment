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
	// make the string into each words
    words := strings.Fields(s)
    for i := range words{
		//make the words title words
      words[i]=strings.Title(words[i])
    }
	//then join each words using join
  camel:= strings.Join(words,"")
  return camel
}

func main(){
  st:="hello world"
  fmt.Println("original : ",st)
  fmt.Println("camelcase : ",CamelCase(st))
}