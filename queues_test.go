package rinqueue

import (
	"testing"
)

type intqueue interface {
	Add(interface{})
	Remove() (interface{}, bool)
	Len() int
	Cap() int
}

func testqueue(t *testing.T, q intqueue) {
	for j := 0; j < 100; j++ {
		if q.Len() != 0 {
			t.Fatal("expected no elements")
		} else if _, ok := q.Remove(); ok {
			t.Fatal("expected no elements")
		}

		for i := 0; i < j; i++ {
			q.Add(i)
		}

		for i := 0; i < j; i++ {
			if x, ok := q.Remove(); !ok {
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
			q.Add(a)
			a++
		}

		for i := 0; i < 2; i++ {
			if x, ok := q.Remove(); !ok {
				t.Fatal("expected an element")
			} else if x != r {
				t.Fatalf("expected %d got %d", r, x)
			}
			r++
		}
	}

	if q.Len() != 200 {
		t.Fatalf("expected 200 elements have %d", q.Len())
	}
}

func TestSlicequeue(t *testing.T) {
	testqueue(t, NewSlicequeue())
}

func TestRingqueue(t *testing.T) {
	testqueue(t, NewRingqueue())
}

func benchmarkAdd(b *testing.B, q intqueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.Add(i)
	}
}

func BenchmarkSliceAdd(b *testing.B) {
	benchmarkAdd(b, NewSlicequeue())
}

func BenchmarkRingAdd(b *testing.B) {
	benchmarkAdd(b, NewRingqueue())
}

func benchmarkRemove(b *testing.B, q intqueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.Add(i)

		if q.Len() > 10 {
			q.Remove()
		}
	}
}

func BenchmarkSliceRemove(b *testing.B) {
	benchmarkRemove(b, NewSlicequeue())
}

func BenchmarkRingRemove(b *testing.B) {
	benchmarkRemove(b, NewRingqueue())
}
