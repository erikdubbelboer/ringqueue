package rinqueue

type Slicequeue []interface{}

func NewSlicequeue() *Slicequeue {
	q := make(Slicequeue, 0)
	return &q
}

func (q *Slicequeue) Add(i interface{}) {
	*q = append(*q, i)
}

func (q *Slicequeue) Remove() (interface{}, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		i := (*q)[0]
		*q = (*q)[1:]

		if n := cap(*q) / 2; len(*q) <= n {
			nodes := make([]interface{}, len(*q), n)
			copy(nodes, *q)
			*q = nodes
		}

		return i, true
	}
}

func (q *Slicequeue) Cap() int {
	return cap(*q)
}

func (q *Slicequeue) Len() int {
	return len(*q)
}
