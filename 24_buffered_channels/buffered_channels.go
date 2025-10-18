package main

import (
	"fmt"
	"time"
)

// we can send limited size of data without blocking
// this channel can hold a few messages temporarily even if no one is receiving right now.
// we can ignore deadlock using buffred channel means we have limited size to paaa the data without blocking
// buffer size is 0 contains 100 bytes 1 contains 108 kb,
// closing buffered channel is so important otherwise it will go in deadlock after all execution
// real world example is queue system like email sending

func emailSender(emailChan <-chan string, done chan bool) { //<-chan means only can recieve the message
	defer func() { done <- true }() //inline function
	for email := range emailChan {
		fmt.Println("sending email to:", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100) // here 100 is the size
	emailChan <- "1@example.com"
	emailChan <- "2@example.com"

	// rememeber in unbuffered channel this conditon was deadloack because we block at above line so we cant reach to the recieve
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	emailChanBulk := make(chan string, 100) //buffered channel
	done := make(chan bool)                 //for wait_group

	go emailSender(emailChanBulk, done) //creating go routine

	//this will not block now i mean not wait for recieve because of buffer its process the sending and then email sender works
	for i := 0; i < 10; i++ {
		emailChanBulk <- fmt.Sprintf("%d@gmail.com", i) //sprintf function use to format the string emaill now: 1@gmail.com

	}
	//print emediatly because it will in go routine now and channel is not blocking because of buffer
	fmt.Println("done sending!")

	close(emailChanBulk) //closing the channel -> preventing deadlock becaude range in email sender will loop infinite and done always will wait

	<-done //should not exit till we recieve done, and it will come only and only when all operations execute in emailSender
}
