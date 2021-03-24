package main

import "fmt"

func main() {
	// 明示的な定義
	// 初期値
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T", arr1)

	// 値を持たせる
	var arr2 [3]int = [3]int{100, 200}
	fmt.Println(arr2)

	// 暗黙的な定義
	arr3 := [3]string{"A", "B", "C"}
	fmt.Println(arr3)

	// 要素数の省略
	arr4 := [...]string{"C", "D"} // [...]要素の数を自動的に数えてくれる
	fmt.Println(arr4)

	// 要素の代入
	arr2[2] = 300
	fmt.Println(arr2)

	// 配列の要素数を調べる
	fmt.Println(len(arr1))

}
