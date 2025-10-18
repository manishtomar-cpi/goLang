package main

import (
	"fmt"
	"sync"
)

/*
- Use mutux for preventing race consition
- When one or more process modifing one resource so that the modification is not atomic
- Race condition means when multiple processed wants to modify single resource then processed will override one anothee changes
- We can use mutux for this , means we can lock the resource when any process holding it, means one go routine can modify it in single time next one will wait untill
- Downside of mutux is at the point "p.view++" go routine has to wait to unlock that breaches consurrency
- Best practice is to lock only the line or lines which need to modify only not in the full function or logic lock only on modification part
*/
type post struct {
	view int
	mu   sync.Mutex //adding mutux
}

func (p *post) inc(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock() //unlock
		wg.Done()
	}() //inline function of defer why adding unlock? because if some error happnes in view++ compilar never reach unlock(next line)

	p.mu.Lock() //adding lock
	p.view++
	// p.mu.Unlock() //unlock

}

func main() {
	var wg sync.WaitGroup
	myPost := post{view: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}
	wg.Wait()
	fmt.Println(myPost.view) //98,97,100,97,96 not proper 100 because many go routines want to modify the view single time calles race condition.

}
