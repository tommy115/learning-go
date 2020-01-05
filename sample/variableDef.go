package sample

import (
	"fmt"
)

func VariableSample() {
	// n を int で定義
	var n int

	// 複数変数を同じ型で定義
	var x, y, z int

	// まとめて複数定義
	var (
		a, b float64
		name string
	)

	n = 20
	x, y, z = 2, 3, 4

	/*
		n = "hello" // ←コンパイルエラー
	*/
	fmt.Printf("n=%d\n", n)
	fmt.Printf("x=%d, y=%d, z=%d", x, y, z)

	a, b = 0.3, -0.4
	fmt.Printf("a=%f, b=%f\n", a, b)

	name = "これはテストです"
	fmt.Println(name)

	i := 3
	j := true
	f := 3.14
	s := "abc"
	fmt.Printf("i=%T, j=%T, f=%T, s=%T\n", i, j, f, s)

	m := one()
	fmt.Printf("m=%T\n", m)
}

func one() int {
	return 1
}
