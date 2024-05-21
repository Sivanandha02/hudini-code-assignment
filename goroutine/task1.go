package main
// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.

func msgToChannel(c chan string){
	msg:="Hello, World!"
	//message assigned to channel
	c<-msg
}
func main()  {
	//creating channel mesg
	mesg:=make(chan string)
	go msgToChannel(mesg)
	//assigning mesg to newmsg
	newMsg:= <-mesg
	fmt.Println(newMsg)
	
}