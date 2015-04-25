package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Tick(time.Second)
	q := newringqueue()
	//q := newslicequeue()

	for i := 1; i > 0; i++ {
		q.add(i)

		if q.len() > 10 {
			q.remove()
		}

		select {
		case <-t:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("cap: %d, len: %d, used: %d\n", q.cap(), q.len(), m.Alloc)
		default:
		}
	}
}
