package main

import (
	"testing"
)

type intqueue interface {
	add(int)
	remove() (int, bool)
	len() int
	cap() int
}

func testqueue(t *testing.T, q intqueue) {
	for j := 0; j < 100; j++ {
		if q.len() != 0 {
			t.Fatal("expected no elements")
		} else if _, ok := q.remove(); ok {
			t.Fatal("expected no elements")
		}

		for i := 0; i < j; i++ {
			q.add(i)
		}

		for i := 0; i < j; i++ {
			if x, ok := q.remove(); !ok {
				t.Fatal("expected an element")
			} else if x != i {
				t.Fatalf("expected %d got %d", i, x)
			}
		}
	}

	a := 0
	r := 0
	for j := 0; j < 100; j++ {
		for i := 0; i < 4; i++ {
			q.add(a)
			a++
		}

		for i := 0; i < 2; i++ {
			if x, ok := q.remove(); !ok {
				t.Fatal("expected an element")
			} else if x != r {
				t.Fatalf("expected %d got %d", r, x)
			}
			r++
		}
	}

	if q.len() != 200 {
		t.Fatalf("expected 200 elements have %d", q.len())
	}
}

func TestSlicequeue(t *testing.T) {
	testqueue(t, newslicequeue())
}

func TestRingqueue(t *testing.T) {
	testqueue(t, newringqueue())
}

func benchmarkAdd(b *testing.B, q intqueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.add(i)
	}
}

func BenchmarkSliceAdd(b *testing.B) {
	benchmarkAdd(b, newslicequeue())
}

func BenchmarkRingAdd(b *testing.B) {
	benchmarkAdd(b, newringqueue())
}

func benchmarkRemove(b *testing.B, q intqueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.add(i)

		if q.len() > 10 {
			q.remove()
		}
	}
}

func BenchmarkSliceRemove(b *testing.B) {
	benchmarkRemove(b, newslicequeue())
}

func BenchmarkRingRemove(b *testing.B) {
	benchmarkRemove(b, newringqueue())
}
