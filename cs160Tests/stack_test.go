package cs160tests

import (
	"src/cs160"
	"strconv"
	"testing"
)

func TestStack_PushSub(t *testing.T) {
	t.Run("Pushing To Stack", func(t *testing.T) {
		var myStack cs160.Stack[string]

		if myStack.IsEmpty() != true {
			t.Error("Expected: Empty Stack=true; Actual: false")
		}

		myStack.Push("aaa")
		myStack.Push("bbb")
		myStack.Push("ccc")

		if myStack.Size() != 3 {
			t.Error("Expected: Size of 3; Actual: " + strconv.Itoa(myStack.Size()))
		}

		myStack.Push("xxx")
		myStack.Push("yyy")
		myStack.Push("zzz")

		if myStack.Size() != 6 {
			t.Error("Expected: Size of 6; Actual: " + strconv.Itoa(myStack.Size()))
		}
	})
}

func TestStack_PopSub(t *testing.T) {
	t.Run("Popping From Stack", func(t *testing.T) {
		var myStack cs160.Stack[string]

		if myStack.IsEmpty() != true {
			t.Error("Expected: Empty Stack=true; Actual: false")
		}

		myStack.Push("aaa")
		myStack.Push("bbb")
		myStack.Push("ccc")

		if myStack.Pops() != "ccc" {
			t.Error("Expected: ccc")
		}

		if myStack.Pops() != "bbb" {
			t.Error("Expected: bbb")
		}

		if myStack.Size() != 1 {
			t.Error("Expected Size: 1")
		}

		if myStack.Pops() != "aaa" {
			t.Error("Expected: aaa")
		}

		if myStack.Size() != 0 {
			t.Error("Expected Size: 0")
		}
	})
}
