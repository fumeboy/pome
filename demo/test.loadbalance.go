package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func testLB() {
	var count int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			if testDo(int32(i)) == -1 {
				atomic.AddInt32(&count, 1)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println(count)
}
