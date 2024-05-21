package main
import "math"
func PowersOfTwo(n int) []uint64 {
  list:=[]uint64{}
  for i:=0;i<=n;i++{
    res:=math.Pow(2,float64(i))
    list=append(list,uint64(res))
  }
  return list
  // your code goes here
}