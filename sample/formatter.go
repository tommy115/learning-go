package sample

import (
	"fmt"
)

// NumberSample 数値を表示するサンプル
func NumberSample() {
	fmt.Printf("数値=%d\n", 5)
	// 表示書式いろいろ
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
	// パラメータが足りないと、、、
	fmt.Printf("%d年%d月%d日\n", 2020, 1)
	// 過剰なパラメータの場合、、、、 lintで分かるが、、
	fmt.Printf("%d年%d月%d日\n", 2020, 1, 5, 30)
	fmt.Printf("\n")

	// %v を試す
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})

	// %#v を試す
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})

	// %Tは、、、
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})
}
