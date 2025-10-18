package main

import (
	"fmt"
	"sync"
)

/*
Use to synchronize our go routines, like we used time.sleep in go that would not be the best practice because we dont know how much time it will take to complete, so we use wait groups in this situation, helps us also in assureance that all go routines are completed

We can use waitgroup from the sync package -> Go provides a WaitGroup from the sync package — it’s a clean way to wait until all goroutines are done.

# Defer use basically as destructor what we wrote in defer will run when all the exuctions of the function is done mean at last

Create a WaitGroup (var wg sync.WaitGroup) -> Add a task before starting goroutine (wg.Add(1)) -> Run goroutine (go task(i, &wg)) -> Inside goroutine, mark task done at end (defer w.Done()) -> Wait in main until all done (wg.Wait())
*/
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // This decreases the counter by 1.
	//When the 1st goroutine finishes → counter = 9 When the 2nd finishes → counter = 8 … and so on. When the counter reaches 0, that means all goroutines are done.
	fmt.Println("doing task", id)
}
func main() {

	//creating waitgroup
	var wg sync.WaitGroup // This creates a WaitGroup object.
	for i := 0; i <= 10; i++ {
		//Before we start a goroutine, we tell the WaitGroup that you’re adding one more task, This increases an internal counter by1.
		wg.Add(1)
		go task(i, &wg)
	}
	//The counter inside wg becomes 10.

	//This line blocks the main goroutine until all other goroutines finish their work (i.e., counter = 0).
	wg.Wait()
}
