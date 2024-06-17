package main

import (
	"fmt"
	"math"
)
const PI = 3.141592653589793; // Use this value as pi in your calculations if necessary

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
  // your code here
  signValue := math.Sin((2*PI)/float64(numberOfSides))
  radiusSqaure := circleRadius*circleRadius
  areaOfPolygon := (float64(numberOfSides)*radiusSqaure/2)*signValue
  return math.Round(areaOfPolygon*1000)/1000
}
func main(){
	re := AreaOfPolygonInsideCircle(4,5)
	fmt.Print("Enter the radius of the circle: ",re)
}