package cs160

/* 
The Stack Type that represents the a traditional FIFO stack
*/
type Stack[T any] struct {
	arrayStack []T
}

/* 
Pushes an item to the top of the stack.
*/
func (stack *Stack[T]) Push(item T) {

	// Modified the array by adding to the array and increasing the length
	stack.arrayStack  = append(stack.arrayStack, item)
}

/* 
Removes the item at the top of the list 
*/
func (stack *Stack[T]) Pops() T {

	// If the length of the stack is zero, we give a null T item
	if stack.Size() == 0 {
		var zero T
		return zero
	}

	// Get the item to be popped and slice it from the array
	poppedItem := stack.arrayStack[stack.Size() - 1]
	stack.arrayStack = stack.arrayStack[:stack.Size() - 1]

	// Return the popped item from array
	return poppedItem
}

/* 
Gets the current size of the stack
*/
func (stack *Stack[T]) Size() int {
	return len(stack.arrayStack)
}

/* 
Boolean value for whether the current stack is empty or not
*/
func (stack *Stack[T]) IsEmpty() bool {
	return stack.Size() == 0
}
