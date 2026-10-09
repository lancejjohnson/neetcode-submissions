type Deque struct {
	values []int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) IsEmpty() bool {
	return len(d.values) == 0
}

func (d *Deque) Append(value int) {
	d.values = append(d.values, value)
}

func (d *Deque) AppendLeft(value int) {
	d.values = append([]int{value}, d.values...)
}

func (d *Deque) Pop() int {
	if len(d.values) == 0 {
		return -1
	}
	last := len(d.values) - 1
	v := d.values[last]
	d.values = d.values[:last]
	return v
}

func (d *Deque) PopLeft() int {
	if len(d.values) == 0 {
		return -1
	}
	v := d.values[0]
	if len(d.values) == 1 {
		d.values = d.values[:0]
	} else {
		d.values = d.values[1:]
	}
	return v
}

