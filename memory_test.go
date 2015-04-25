package main

import (
	"math"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func benchmarkMemory(b *testing.B, q intqueue) {
	b.SkipNow()
	b.ReportAllocs()

	b.N = 30000000

	for i := 0; i < b.N; i++ {
		q.add(i)
	}

	for i := 0; i < b.N; i++ {
		q.remove()
	}

	b.Logf(memory())
}

func BenchmarkSliceMemory(b *testing.B) {
	benchmarkMemory(b, newslicequeue())
}

func BenchmarkRingMemory(b *testing.B) {
	benchmarkMemory(b, newringqueue())
}

func memory() string {
	runtime.GC()
	time.Sleep(time.Second)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	units := []string{
		"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB",
	}

	base := int(math.Floor(math.Log(float64(m.Alloc)) / math.Log(1024)))
	return strconv.FormatFloat(float64(m.Alloc)/math.Pow(1024, float64(base)), 'f', 2, 64) + " " + units[base]
}
