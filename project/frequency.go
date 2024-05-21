package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//create the function that give the frquency of the words
func frequencyOfWord(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text:")
	text, _ := reader.ReadString('\n')
	//fmt.Println("text is ",text)
	
	//create a map for storing the frequencies of the words
	freqStore:=map[string]int{}
	//split the text into word
	words:=strings.Fields(text)
	//iterate the words and increment the frquencies of words
	for _,word:= range words{
		freqStore[word]++
	}
	// for i:=0;i<len(words);i++{
	// 	freqStore[words[i]]++
	// }
	//iterate the map freqstore and display the word and frequency
	for w,f:=range freqStore{
		fmt.Printf("'%s' - %d\n",w,f)
	}
}



func main()  {
	frequencyOfWord()
}