//rw mutex

package main

import (
"fmt"
"sync"
"time"
)

var (
rwMutex sync.RWMutex
data    = 0
)

func reader(id int, wg *sync.WaitGroup) {
rwMutex.RLock()
fmt.Printf("Reader %d: read data = %d\n", id, data)
time.Sleep(100 * time.Millisecond)
rwMutex.RUnlock()
wg.Done()
}

func writer(id int, wg *sync.WaitGroup) {
rwMutex.Lock()
data += 1
fmt.Printf("Writer %d: wrote data = %d\n", id, data)
time.Sleep(150 * time.Millisecond)
rwMutex.Unlock()
wg.Done()
}

func main() {
var wg sync.WaitGroup

for i := range(4) {
wg.Add(1)
go reader(i, &wg)
}

for i := 1; i <= 2; i++ {
wg.Add(1)
go writer(i, &wg)
}

wg.Wait()
}

