package main

import (
	"fmt"
)

type NestedInteger struct {
	Data interface{}
	List []*NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return n.List != nil
}

func (n NestedInteger) GetInteger() int {
	val, ok := n.Data.(int)

	if ok {
		return val
	}

	return 0
}

func (n NestedInteger) GetList() []*NestedInteger {
	if n.List == nil {
		return make([]*NestedInteger, 0)
	}

	return n.List
}

func calculate(nestedList []*NestedInteger, depth int) int {
	result := 0

	for _, item := range nestedList {
		if item.IsInteger() {
			result += item.GetInteger() * depth
		} else {
			result += calculate(item.GetList(), depth+1)
		}
	}

	return result
}

func depthSum(nestedList []*NestedInteger) int {
	return calculate(nestedList, 1)
}

func main() {
	daList := []*NestedInteger{&NestedInteger{Data: 42, List: []*NestedInteger{&NestedInteger{Data: 42, List: nil}}}}
	fmt.Println(
		depthSum(daList),
	)
}
