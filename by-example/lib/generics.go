package lib

// a generic find-index-of-target-in-slice
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// generic function
func MapKeys[K comparable, V any](amap map[K]V) []K {
	keys := make([]K, len(amap))
	var count int
	for key := range amap {
		keys[count] = key
		count++
	}
	return keys
}

// generic type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}

func (lst *List[T]) Add(value T) {
	newElement := &element[T]{val: value}
	if lst.tail == nil {
		lst.head = newElement
		lst.tail = lst.head
	} else {
		lst.tail.next = newElement
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) ToSlice() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
