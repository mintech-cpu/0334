package main

import (
	"fmt"
	"strings"
)

func main() {
	// 3 + 1 の答え
	fmt.Println(3 + 1)
	// 実行結果が「3 + 1 = 4」
	fmt.Println("3 + 1 =", 3+1)
	// 3 - 1の答え
	fmt.Println(3 - 1)
	// 実行結果が「3 - 1 = 2」
	fmt.Println("3 - 1 =", 3-1)
	// 3 × 1の答え
	fmt.Println(3 * 1)
	// 実行結果が「3 × 1 = 3」
	fmt.Println("3 × 1 =", 3*1)
	// 3 / 1の答え
	fmt.Println(3 / 1)
	// 実行結果が「3 / 1 = 3」
	fmt.Println("3 / 1 =", 3/1)
	// 30を4で割った時の余り
	fmt.Println(30 % 4)

	// 「Hello Go」を表示(三通り)
	fmt.Println("Hello Go")
	fmt.Println("Hello " + "Go")
	s1 := "Hello "
	s1 += "Go"
	fmt.Println(s1)

	// 「Hello Go」2番目の文字「e」を表示
	fmt.Println(string("Hello Go"[1]))
	// 「Hello Go」を「Hello go」に変更
	s2 := "Hello Go"
	s2 = strings.Replace(s2, "G", "g", 1)
	fmt.Println(s2)
	// =>sに代入されている「Hello Go」の"G"を"g"に置き換える。最初の一個だけ
	// sの文字列に「Hello」が含まれているか？ => true or false
	fmt.Println(strings.Contains(s2, "Hello"))

	// trueとfalseの型を表示
	t, f := true, false
	fmt.Printf("%T", t)
	fmt.Printf("%T", f)

	// 論理演算子
	// 実行結果を答えてください
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true || false == true)
	fmt.Println(false || false == true)
	fmt.Println(!false) // 論理値の反転

}
