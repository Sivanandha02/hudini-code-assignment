package main

// 1. Calculate average
//declard the array with datatype float
var num = []float64{1, 2, 3, 4}

// create the function the return type is float
func calculateAverage() float64 {
	//declared sum as zero with float
	var sum float64 = 0
	//iterate each value in array
	for i := 0; i < len(num); i++ {
		//add each element in the array
		sum += num[i]
	}
	//then return the average
	return sum / float64(len(num))
}

// 2. Check Age
func checkAge(age int) string {
	//create switch for compare the conditions
	switch {
	case age < 16 && age >= 1:
		return "minor"
	case age >= 16 && age <= 40:
		return "young adult"
	case age > 40:
		return "adult"
	default:
		return "enter the correct value"
	}
}

// 3. Reverse of string
func reverseString(str string) string {
	var revstr string
	for i := len(str) - 1; i >= 0; i-- {
		revstr += string(str[i])
	}
	return revstr
}

// 4. FInd the Largest
var numbers = []int{2, 4, 3, 7, 5, 8, 1}

func largestNumber(numbers []int) int {
	var large int = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if large < numbers[i] {
			large = numbers[i]
		}
	}
	return large
}

//5. Counter
// created the counter struct it is userdefined datatype
type Counter struct{
    //initialize the property
    countn int
}
//create add function  like increment the parameter
func add(a Counter) Counter{
    a.countn++
    return a
}
//create decrement function
func decrement(a Counter) Counter {
    a.countn--
    return a
}
//create reset function to reset the parameter back to 0
func reset(a Counter) Counter{
    a.countn=0
    return a
}
