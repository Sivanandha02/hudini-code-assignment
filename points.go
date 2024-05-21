package main
import "strings"
func Points(games []string) int {
  // your code here!
  score:=0
  for i:=0;i<len(games);i++{
    res:=strings.Split(games[i],":")
    if res[0]>res[1]{
    score+=3
  }else if res[0]<res[1]{
    score+=0
  }else{
    score+=1
  }
  }
  
  return score
}