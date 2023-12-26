package cs160tests

import (
	"src/cs160"
	"testing"
)


func TestDeque_FrontSub(t *testing.T) {
	t.Run("Pushing/Popping To Front Of Deque", func(t *testing.T) {
		var myDeque cs160.Deque[string]

		if myDeque.IsEmpty() != true {
			t.Error("Expected: EmptyDeque=true; Actual: false")
		}

		myDeque.AssignArray([]string{"soren", "vincent", "cam"})

		if myDeque.IsEmpty() != false {
			t.Errorf("Expected: IsEmpty=false; Actual: %v", myDeque.IsEmpty())
		}

		myDeque.PushFront("christian")
		myDeque.PushFront("michael")
		myDeque.PushFront("noah")

		if myDeque.Size() != 6 {
			t.Errorf("Expected: Size=6; Actual: %d", myDeque.Size())
		}

		popped_item1 := *myDeque.PopFront()

		if popped_item1 != "noah" {
			t.Errorf("Expected Popped='noah'; Actual: %s", popped_item1)
		}

		popped_item2 := *myDeque.PopFront()

		if popped_item2 != "michael" {
			t.Errorf("Expected Pop='michael'; Actual: %s", popped_item2)
		}

		if myDeque.Size() != 4 {
			t.Errorf("Expected size of 4; Actual: %d", myDeque.Size())
		}

	})
}

func TestDeque_BackSub(t *testing.T) {
	t.Run("Pushing/Popping To Back Of Deque", func(t *testing.T) {
		var myDeque cs160.Deque[string]

		if myDeque.IsEmpty() != true {
			t.Error("Expected: EmptyDeque=true; Actual: false")
		}

		myDeque.AssignArray([]string{"soren", "vincent", "cam"})

		if myDeque.IsEmpty() != false {
			t.Errorf("Expected: IsEmpty=false; Actual: %v", myDeque.IsEmpty())
		}

		myDeque.PushBack("christian")
		myDeque.PushBack("michael")
		myDeque.PushBack("noah")

		if myDeque.Size() != 6 {
			t.Errorf("Expected: Size=6; Actual: %d", myDeque.Size())
		}

		popped_item1 := *myDeque.PopBack()

		if popped_item1 != "noah" {
			t.Errorf("Expected Popped='noah'; Actual: %s", popped_item1)
		}

		popped_item2 := *myDeque.PopBack()

		if popped_item2 != "michael" {
			t.Errorf("Expected Pop='michael'; Actual: %s", popped_item2)
		}

		if myDeque.Size() != 4 {
			t.Errorf("Expected size of 4; Actual: %d", myDeque.Size())
		}

	})
}