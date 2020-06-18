package main

import (
	"fmt"
)

func main() {
	arith()
	arrays()
	slices()
	maps()
	loops()
}

func arith() {
	x := 5 //shorthand implemntation
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)
}

func arrays() {
	var array [5]int
	fmt.Println(array)

	arrayShorthand := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayShorthand)
}

func slices() {
	slice := []int{5, 4, 3, 2, 1}
	fmt.Println(slice)
	slice = append(slice, 0) //this creates a new slice
	fmt.Println(slice)
}

func maps() {
	vertices := make(map[string]int)
	vertices ["triangle"] = 3
	vertices ["square"] = 4
	vertices ["pentagon"] = 5
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	delete(vertices, "pentagon")
	fmt.Println(vertices)
}

func loops() {
	//normal for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loop style
	j := 0 
	for j < 5 {
		fmt.Println(j)
		j++
	}
}