package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		//assert stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		//add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		//add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())

		//can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		//assert stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		//add a thing, check it's not empty
		myStackOfStrings.Push("abc")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		//add another thing, then pop it back again
		myStackOfStrings.Push("def")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "def")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "abc")
		AssertTrue(t, myStackOfStrings.IsEmpty())

		//can get the numbers we put in as numbers, not untyped interface{}
		myStackOfStrings.Push("hello")
		myStackOfStrings.Push("world")
		world, _ := myStackOfStrings.Pop()
		hello, _ := myStackOfStrings.Pop()
		AssertEqual(t, hello+" "+world, "hello world")
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want%v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("got %v, want%v", got, want)
	}
}
