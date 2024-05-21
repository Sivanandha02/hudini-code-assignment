package main
import "strings"
func FindShort(s string) string {
  short:=[] string{}
  short=strings.Split(s," ")
  min:=short[0]
  for i:=1;i<len(short);i++{
    if len(min)>len(short[i]){
      min=short[i]
    }
  }
  return min
  // your code
}