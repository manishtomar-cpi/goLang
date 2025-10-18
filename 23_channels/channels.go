package main

import (
	"fmt"
)

/*
- Working as a bridge bw go routines , like we want to pass the data bw go routines we can pass with the help of channels
- Channel is blocking till second side of the pipe is not ready to recieve the data so deadlock occure in this condition when no other side channel found to recieve the data
- We can use channel as a queue in our program
- We can do visa versa also who is recieving can send and who is sending can recieve
- We can do same work from channel that we did from wait group
- We use done in difer because sometime program cash in processing the not come to done if we use done at last so after crash done should execute

unbuffered channel -> un buffered channel is blocking means program should block at the stage, one by one communication, means we are sending and waiting, means we can send one data at the time we should wait for recieve then we can send next data, slow in the process so we used buffered channel,

Why single unbuffered channel goes into deadlock?
The main function itself is a goroutine. When ch <- "ping" runs, it waits for someone to receive. But there’s no other goroutine running that can receive. So the main goroutine blocks forever. If still no receiver ever appears, Go realizes nothing can continue error - deadlock!
*/

func processNum(numChan chan int) {
	fmt.Println("processing number", <-numChan) // we are recieving message from channel by <- numchan
}

func processNumSecond(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num) // we are recieving message from channel by <- numchan
	}
}

// send data from channel
func sum(result chan int, num1 int, num2 int) {
	addition := num1 + num2
	result <- addition // now channel sending the data
}

// channel as wait group
func task(done chan bool) {
	defer func() { done <- true }() //creating in line function in defer which run at last when function execute
	fmt.Println("processing...")
}
func main() {

	//making channel

	// messageChan := make(chan string)
	// messageChan <- "ping" //sending message to channel -> This line blocks — it will wait until some other goroutine is ready to receive the message.
	// msg := <-messageChan  //getting data from channel

	// fmt.Println(msg)

	numChan := make(chan int) //creating channel

	go processNum(numChan) //making go routine
	numChan <- 5           // sending the data

	numChanSecond := make(chan int)
	go processNumSecond(numChanSecond)

	fmt.Println("second channel...")

	// for {
	// 	numChanSecond <- rand.Intn(100)
	// 	time.Sleep(time.Second)
	// }

	//recieving the data
	sumChan := make(chan int)
	go sum(sumChan, 2, 3)
	fmt.Println(<-sumChan) //recieving the message

	//channel as waitgroup
	done := make(chan bool)
	go task(done)

	<-done // this will block the program once we get the message then the program should exit so that is same as wait group

	//handling multiple chanlles
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 1
	}()
	go func() {
		chan2 <- "pong"
	}()

	//use select case -> select case only use in channels
	//It waits for multiple communication operations and executes whichever one is ready first.
	//select is for channels — it’s how Go handles concurrent communication.
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recieved data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("recieved data from chan2", chan2Val)
		}
	}

}
