package main

import (
	"fmt"
	"runtime"
	"time"

	rq "github.com/ErikDubbelboer/ringqueue"
)

func main() {
	t := time.Tick(time.Second)
	//q := rq.NewRingqueue()
	q := rq.NewSlicequeue()

	for i := 1; i > 0; i++ {
		q.Add(i)

		if q.Len() > 10 {
			q.Remove()
		}

		select {
		case <-t:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("cap: %d, len: %d, used: %d\n", q.Cap(), q.Len(), m.Alloc)
		default:
		}
	}
}
