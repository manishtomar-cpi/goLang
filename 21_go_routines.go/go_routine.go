package main

import (
	"fmt"
	"time"
)

// Light weight threads, use to do multithreading, achieve concurency
// We can use in crone jobs background jobs

func task(id int) {
	fmt.Println("doing task", id)
}
func main() {
	for i := 0; i <= 10; i++ {
		// task(i) //doing task 0, doing task 1, doing task 2 ...... because every time function call running one by one like 0-10

		//if we want to run these threads parellel we can use:
		go task(i)
		//doing task 10, doing task 2, doing task 4, doing task 7, doing task 8, doing task 6, doing task 9, doing task 5 doing task 3, doing task 0, doing task 1 -> now these all tasks happening in IIel thats why order changed
	}
	// fmt.Println("main done")
	time.Sleep(time.Second) // use because when our go routines was unning our program is exit

}
