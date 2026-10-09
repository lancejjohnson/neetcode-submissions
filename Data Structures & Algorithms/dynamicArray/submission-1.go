type DynamicArray struct {
	values []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{values: make([]int, 0, capacity)}
}

func (a *DynamicArray) GetSize() int {
	return len(a.values)
}

func (a *DynamicArray) GetCapacity() int {
	return cap(a.values)
}

func (a *DynamicArray) Get(i int) int {
	return a.values[i]
}

func (a *DynamicArray) Set(i, n int) {
	a.values[i] = n
}

func (a *DynamicArray) Popback() int {
	v := a.values[len(a.values)-1]
	a.values = a.values[:len(a.values)-1]
	return v
}

func (a *DynamicArray) Pushback(value int) {
	if a.GetSize() >= a.GetCapacity() {
		a.resize()
	}
	a.values = append(a.values, value)
}

func (a *DynamicArray) resize() {
	newValues := make([]int, 0, a.GetCapacity()*2)
	for _, v := range a.values {
		newValues = append(newValues, v)
	}
	a.values = newValues
}

