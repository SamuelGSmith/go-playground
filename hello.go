package main

import (
	"fmt"
	"errors"
	"math"
)

//creating a struct, which is a collection of fields group to create a more logical type

type person struct {
	name string
	age int
}

func main() {
	arith()
	arrays()
	slices()
	maps()
	loops()
	//square root - 'happy case'
	resultHappy, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultHappy)
	}

	//square root - 'unhappy case'
	resultUnhappy, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultUnhappy)
	}

	createPerson("Sam", 26)
	pointers()
}

// function format:
// func name(arg type, arg type) return-type

// functions can have multiple return values in GO
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

	//array loop
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	//map loop
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range arr {
		fmt.Println("key:", key, "value:", value)
	}	
}

func sqrt(x float64) (float64, error){
	if x < 0 {
		//errors need to be returned in this manner, as GO does not have any exceptions
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}	

func createPerson (name string, age int) {
	p := person{name: name, age: age}
	fmt.Println(p)
	//print specific field
	fmt.Println(p.age)
}

func pointers(){
	i := 7
	fmt.Println(i)
	//the ampersand gives us the memory location
	inc(&i)
	fmt.Println(i)

}

func inc (x *int) {
	*x++
}