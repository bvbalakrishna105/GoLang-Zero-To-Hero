package main

import (
	"errors"
	"fmt"
	"os"
)

func describe(i interface{}) {
	switch v := i.(type) {

	case int:
		fmt.Println("Type : int")
	case string:
		fmt.Println("Type is :  string")
	default:
		fmt.Printf("Unknow type %v", v)

	}

}

func main() {
	//describe(42)
	//describe("Hello")
	//arraySlice()
	//transposeMatrix(
	//singleLinkedList()

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Result :", result)
	}

	/*fmt.Println("start")
	recoverExample()
	fmt.Println("End")*/
	fileErrorAssertion()
}

func arraySlice() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println(arr[2])
	fmt.Println(slice[0])

	newSlice := []int{1, 2, 3, 4}
	index := 2
	value := 10

	newSlice = append(newSlice[:index], append([]int{value}, newSlice...)...)
	fmt.Println(newSlice)
}

func transposeMatrix() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var resultMatrix [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			resultMatrix[i][j] = matrix[j][i]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", resultMatrix[i][j])
		}
		fmt.Printf("\n")
	}

}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data int) {

	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func singleLinkedList() {

	list := LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("Something went wrong")
}

func fileErrorAssertion() {
	file, err := os.Open("text.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error: ", err)
		}
	} else {
		defer file.Close()
		fmt.Println("File Opened sucessfully")
	}
}
