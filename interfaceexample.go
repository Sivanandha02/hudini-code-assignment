package main
import "fmt"
func main(){
	var runner Animal
	runner=dog{name: "dog1"}
	runner.Run()
}

type Animal interface{
	Run()
}
type dog struct{
	name string
}
func (d dog) Run(){
	fmt.Println(d.name,"is running")
}