package main

// ポインタQ4.関数を用いて、変数numに5をかけてください。

import "fmt"

func Add(i *int) {
	*i = *i * 5

}

func main() {

	// Q1.値に100を持つ変数nを定義して、nのアドレスを表示してください。
	var n int = 100
	fmt.Println(&n)

	// Q2.値をnのアドレスとしたポインタ型変数pを定義し、pのアドレスとpの値を表示してください。
	var p *int = &n
	fmt.Println(p)  // pのアドレス
	fmt.Println(*p) // pの値

	// Q3.関数を用いずに、変数nとpの値を同時に200へ変更し、nとpの値を表示してください。
	n = 200
	fmt.Println(n)
	fmt.Println(*p)

	// Q4.関数を用いて、変数numに5をかけてください。
	var num int = 100
	Add(&num)
	fmt.Println(num)
}
