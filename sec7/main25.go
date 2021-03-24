package main

import (
	"fmt"
)

func Plus(x, y int) int {
	return x + y
}
func Div(x, y int) (int, int) {
	q := x / y
	p := x % y
	return q, p
}

// 返り値を表す変数
func Double(price int) (result int) {
	result = price * 3
	return
}

func main() {
	var i int = Plus(1, 2)
	fmt.Println(i)

	var i2, i3 = Div(6, 2)
	fmt.Println(i2, i3)

	var i4 int = Double(5)
	fmt.Println(i4)

}
