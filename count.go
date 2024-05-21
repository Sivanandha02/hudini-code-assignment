
package main
func CountBy(x, n int) []int {
	list:=[]int{}
	for i:=1;i<=n;i++{
		list=append(list,i*x)
	}
	return list
  }
  