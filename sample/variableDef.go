package sample

import (
	"fmt"
	"math"
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

	// 配列
	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a=%v\n", a1)

	// 配列2 これでもOK
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("a2=%v\n", a2)

	// 配列3 少なめに与えてみると、残りの要素は0になる。
	a3 := [5]int{6, 7, 8}
	fmt.Printf("a3=%v\n", a3)

	// 配列コピー
	a3 = a1
	fmt.Printf("a3=%v\n", a3)

	// a3を一部変更
	a3[2] = 100
	fmt.Printf("a3=%v\n", a3)

	// a1は変わらない
	fmt.Printf("a1=%v\n", a1)

	var ii interface{}
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)

	// 何でも代入可能
	ii = 1
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)
	ii = 3.14
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)
	ii = '山'
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)
	ii = "文字列"
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)
	ii = [...]int{1, 2, 3}
	fmt.Printf("ii=%v, type ii=%T\n", ii, ii)
}

func one() int {
	return 1
}

func ChangeVariable() {
	x := uint(17)
	n := 1
	b := byte(n)
	i64 := int64(n)
	c := 5 + 7i
	r := '松'
	s := "文字列"
	s2 := `文字列
	の生の状態
	です。`

	fmt.Printf("x=%#v, n=%#v, b=%#v, i64=%#v, c=%#v, r=%#v, s=%#v s2=%#v\n", x, n, b, i64, c, r, s, s2)
	fmt.Printf("x=%T, n=%T, b=%T, i64=%T, c=%T, r=%T, s=%T, s2=%T\n", x, n, b, i64, c, r, s, s2)
}

func MathVariable() {
	fmt.Printf("uint32 max value=%d\n", math.MaxUint32)
}
