package main

import (
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		time.Sleep(1 * time.Second)

		mu2.Lock()
		defer mu2.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()
		time.Sleep(1 * time.Second)
		mu1.Lock()
		defer mu1.Unlock()
	}()
	wg.Wait()
}
