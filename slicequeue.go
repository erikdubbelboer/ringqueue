package main

type slicequeue []int

func newslicequeue() *slicequeue {
	q := make(slicequeue, 0)
	return &q
}

func (q *slicequeue) add(i int) {
	*q = append(*q, i)
}

func (q *slicequeue) remove() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		i := (*q)[0]
		*q = (*q)[1:]

		if n := cap(*q) / 2; len(*q) <= n {
			nodes := make([]int, len(*q), n)
			copy(nodes, *q)
			*q = nodes
		}

		return i, true
	}
}

func (q *slicequeue) cap() int {
	return cap(*q)
}

func (q *slicequeue) len() int {
	return len(*q)
}
