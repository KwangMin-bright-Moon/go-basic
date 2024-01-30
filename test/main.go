package main

import (
	"fmt"
	"strings"
)

func add (a int, b int) int {
	return a + b
}


func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string){
	defer fmt.Println("I'm done!")
	defer fmt.Println("Hi")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) []string{
	fmt.Println(words)
	return words
}

func loop (number ...int) {
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}

// func superAdd (number ...int) int {
// 	total := 0;
	
// 	for i := 0; i < len(number); i++ {
// 		total += number[i]
// 	}

// 	return total
// }

func superAdd (number ...int) int {
	total := 0;

	for _, number := range number {
		total += number
	}

	return total
}

func main(){
	result := superAdd(1, 2, 3, 5)
	fmt.Println(result)

	
}
