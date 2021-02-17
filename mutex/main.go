package main

import (
	"fmt"
	"sync"
)

var (
	counter       = 0
	atomicCounter atomicInt
	lock          sync.Mutex
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *atomicInt) increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *atomicInt) derease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *atomicInt) valueOfInt() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
	fmt.Printf("Final atomic counter: %d\n", atomicCounter.value)
}

func updateCounter(wg *sync.WaitGroup) {
	// without the mutex, the final counter wouldn't be 10000
	// Mutex restricts the calls to the counter so one call at
	// a time can be made.
	lock.Lock()
	defer lock.Unlock()

	counter++
	atomicCounter.increase()
	wg.Done()
}
