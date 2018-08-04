package main

import (
	"errors"
	"fmt"
)

const plus = -1
const minus = -2

func main() {

	err := calc([]int{2, 9, plus, 5, minus})
	if err != nil {
		fmt.Println("Your stuff broke.", err)
	}
}

func calc(tokens []int) error {
	var err error
	var stack []int
	for _, token := range tokens {
		switch token {
		case plus:
			var a, b int
			if len(stack) < 2 {
				return errors.New("Less than two on stack.")
			}
			stack, a, err = Pop(stack)
			if err != nil {
				return fmt.Errorf("Tried to pop, did not work: %s", err)
			}
			stack, b, err = Pop(stack)
			if err != nil {
				return fmt.Errorf("Tried to pop, did not work: %s", err)
			}
			stack = Push(stack, a+b)
			fmt.Println(stack)
		case minus:
			var a, b int
			if len(stack) < 2 {
				return errors.New("Less than two on stack.")
			}
			stack, a, err = Pop(stack)
			if err != nil {
				return fmt.Errorf("Tried to pop, did not work: %s", err)
			}
			stack, b, err = Pop(stack)
			if err != nil {
				return fmt.Errorf("Tried to pop, did not work: %s", err)
			}
			stack = Push(stack, b-a)
			fmt.Println(stack)
		default:
			stack = Push(stack, token)
			fmt.Println(stack)

		}

	}
	// Each token-
	//
	//Given number or operator
	//
	//Numbers goven are added to the stack
	//
	//Operator when given are checked to the stack if there is two on top they arepoped off the stack and computed
	//
	//result is pushed on to the stack
	//
	//starting numbers are discarded if not avali error

	return nil
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
