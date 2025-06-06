package main
//maps are not thread-safe in Go, so we need to use a
//mutex to protect access to the shared variable.

 import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	count++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Count:", count)
}

