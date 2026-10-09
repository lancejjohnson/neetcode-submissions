type LinkedList struct {
	head *ListNode
	len  int
}

type ListNode struct {
	value int
	next  *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	node := ll.head

	for i := range index + 1 {
		if node == nil {
			return -1
		}
		if i == index {
			return node.value
		}
		node = node.next
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	ll.head = &ListNode{value: val, next: ll.head}
	ll.len++
}

func (ll *LinkedList) InsertTail(val int) {
	tail := &ListNode{value: val}
	if ll.head == nil {
		ll.head = tail
		ll.len++
		return
	}

	node := ll.head
	for true {
		if node.next == nil {
			node.next = tail
			ll.len++
			return
		}
		node = node.next
	}
}

func (ll *LinkedList) Remove(index int) bool {
	var node, prev *ListNode = ll.head, nil

	for i := range index + 1 {
		if node == nil {
			return false
		}
		if i == index {
			if prev == nil {
				ll.head = node.next
			} else {
				prev.next = node.next
				node.next = nil
			}
			ll.len--
			return true
		}

		prev, node = node, node.next
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, 0, ll.len)
	node := ll.head

	for node != nil {
		values = append(values, node.value)
		node = node.next
	}

	return values
}

