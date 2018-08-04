package main

import (
	"errors"
	"fmt"
)

func main() {

	var poplast int
	var err error
	var stack []int
	stack = Push(stack, 1)
	fmt.Println(stack)

	stack = Push(stack, 2)
	fmt.Println(stack)

	stack = Push(stack, 3)
	fmt.Println(stack)

	stack = Push(stack, 4)
	fmt.Println(stack)

	stack = Push(stack, 5)
	fmt.Println(stack)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)

	stack, poplast, err = Pop(stack)
	fmt.Println(stack, poplast, err)
}

func Push(stack []int, element int) []int {
	//func main() {
	//	aSlice := []int{1,2,3,4,5}
	//	fmt.Println(aSlice)
	//	aSlice=append(aSlice, 1)
	//	fmt.Println(aSlice)
	var newstack []int = append(stack, element)
	return newstack

}
func Pop(stack []int) ([]int, int, error) {

	var numbers []int
	var last int
	var err error
	var allbutlast = len(stack) - 1

	if len(stack) > 0 {
		numbers = stack[:allbutlast]
		last = stack[allbutlast]
		return numbers, last, err
	}

	err = errors.New("Nothing here to find!")
	return numbers, last, err

}
