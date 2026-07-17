package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	counter int
	lock    sync.Mutex
}

func run(c *Counter, wg *sync.WaitGroup) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.counter++
	fmt.Println(c.counter)
	wg.Done()
}

func somethingWithTime(a int, b int, ch chan<- int) {
	time.Sleep(5 * time.Second)

	ch <- a + b
}

func main() {

	c := Counter{counter: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go run(&c, &wg)
	}

	wg.Wait()
	fmt.Println("done tha task")

	// ch := make(chan int)

	// go somethingWithTime(1000000, 100000, ch)

	// x := <-ch
	// fmt.Println(x)

}
