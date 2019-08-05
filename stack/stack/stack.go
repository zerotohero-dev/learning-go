package stack

import (
	"errors"
)

type Stack []interface{}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack

	if len(theStack) == 0 {
		return nil, errors.New("cannot Pop() an empty stack")
	}

	x := theStack[len(theStack)-1]

	*stack = theStack[:len(theStack)-1]

	return x, nil
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack *Stack) Top() (interface{}, error){
	theStack := *stack

	if len(theStack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}

	return theStack[len(theStack)-1], nil
}

func (stack *Stack) Cap() int {
	return cap(*stack)
}

func (stack *Stack) Len() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}
