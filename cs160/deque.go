package cs160

/*
The Deque Type that allows for queue-like operations from 
both ends of the collection
*/
type Deque[T any] struct {
	arrayDeque []T
}

/*
Gets the size of the current deque
*/
func (deque *Deque[T]) Size() int {
	return len(deque.arrayDeque)
}

/*
Given an array, assigns that items in the that array to the deque
*/
func (deque *Deque[T]) AssignArray(array []T) {
	*deque = Deque[T]{arrayDeque: array}
}

/*
Pushes/Adds an item to the front of the deque
*/
func (deque *Deque[T]) PushFront(item T) {
	// Add item to front of deque
	// The '...' spreads the items in the list
	deque.arrayDeque = append([]T{item}, deque.arrayDeque...)
}

/*
Pops/removes an item from the beginning of the deque.
*/
func (deque *Deque[T]) PopFront() *T {
	if deque.Size() == 0 {
		var zero T
		return &zero
	}

	poppedItem := deque.arrayDeque[deque.Size() -1]
	deque.arrayDeque = deque.arrayDeque[:deque.Size()-1]

	return &poppedItem
}
/*
Pushes/Adds an item to the back of the deque
*/
func (deque *Deque[T]) PushBack(item T) {
	// Append an item at the end of the slice
	deque.arrayDeque = append(deque.arrayDeque, item)
}

/* 
Pops the item at the back of the deque, returns the popped item
*/
func (deque *Deque[T]) PopBack() *T {
	if deque.Size() == 0 {
		var zero T
		return &zero
	}

	poppedItem := deque.arrayDeque[0]
	deque.arrayDeque = deque.arrayDeque[0:]

	return &poppedItem
}

/*
Checks if the current deque is empty, if true then deque is empty
*/
func (deque *Deque[T]) IsEmpty() bool {
	return deque.Size() == 0
}