package cs160

type LinkedListNode[T any] struct {
	Value T
	Tail *LinkedListNode[T]
}
