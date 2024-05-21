package main

// func OddCount(n int) int {
// 	//your code here
// 	count := 0
// 	for i := 1; i < n; i++ {
// 		if i%2 != 0 {
// 			count += 1
// 		}
// 	}

// 	return count
// }
func OddCount(n int) int{
	//your code here
	   if n%2!=0{
		return (n-1)/2
	  }else{
		return n/2
	  }
  }
  
