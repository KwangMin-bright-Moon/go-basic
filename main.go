package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) int {
	return a + b
}



func multiply(a,b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string){
// 	return len(name), strings.ToUpper(name)
// }	

func lenAndUpper(name string) (length int, uppercase string){
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func loop (number ...int) {
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}

func superAdd (number ...int) int {
	total := 0;
	for _, number := range number {
		total += number
	}

	return total
}

// func canIDrink(age int) bool {
// 	koreanAge := age + 2
// 	if koreanAge < 18 {
// 		return false
// 	}
// 	return true
// }

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

 	return true
}

// func canIDrive(age int) bool {
// 	switch age {
// 	case 18 :
// 		return false
// 	case 20 :
// 		return true
// 	}
// 	return false
// }

func canIDrive1(age int) bool {
	switch koreanAge := age + 2; koreanAge{
	case 18 :
		return false
	case 20 :
		return true
	}
	return false
}

func canIDrive(age int) bool {
	switch {
	case age < 18 :
		return false
	case age >= 20 :
		return true
	}
	return false
}

func pointer(){
	a := 2
	b := a // 이렇게 하면 값을 복사한다.
	a = 3 
	// 주소를 가리키고 싶을 때 pointer를 쓸 수 있다.
	fmt.Println(a, b)
	fmt.Println(&a, &b) // &를 붙이면 메모리가 출력되는 걸 볼 수 있다.

	// 주소를 가리키게 하고 싶을 경우 아래와 같이 하면 된다.
	c := 2
	d := &c
	c = 3
	fmt.Println(&c, d) // 같은 주소가 출력되는 걸 볼 수 있다.
	// *를 붙이면 주소에 있는 값을 볼 수 있다.
	fmt.Println(*d) // 3이 출력된다. c의 마지막에 넣어준 값이 3이기 때문에
	// d를 업데이트해서 c를 변경할 수도 있다.
	*d = 10
	fmt.Println(c) // 10이 출력된다.

}

func main(){
	pointer()
}