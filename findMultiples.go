package main
func FindMultiples(integer, limit int) []int {
	// Your code here!
	list:=[]int {}
	for i:=1;i<=limit;i++{
	  if i%integer==0{
		list=append(list,i)
	  }
	}
	return list
  }
// for word, freq := range wordFreqs {
	//fmt.Println(word, freq)
