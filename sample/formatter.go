package sample

import (
	"fmt"
)

// NumberSample 数値を表示するサンプル
func NumberSample() {
	fmt.Println("数値=%d\n", 5)
	fmt.Println("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
}
