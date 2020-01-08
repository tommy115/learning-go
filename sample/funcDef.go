package sample

import "fmt"

func init() {
	fmt.Println("init in funcDef")
}

// FuncSample ... 関数サンプル
func FuncSample() {
	fmt.Printf("func1=%d\n", func1(1, 2))
	fmt.Printf("func2=%d\n", func2(3, 4))
	fmt.Printf("func3=%f\n", func3(0.25, 3.4, 3, 4))

	// 複数戻り値を受け取る
	x, y := func4(2, 5)
	fmt.Printf("func4=%d, %d\n", x, y)

	// 片方は受けなくてもよい
	_, b := func4(5, 4)
	fmt.Printf("func4=%d\n", b)

	// 呼び出した後にエラーチェック
	result, err := func5(2, 3)
	if err != nil {
		fmt.Printf("error が起きました%v\n", err)
	} else {
		fmt.Printf("result=%d\n", result)
	}

	// 戻り値に変数名があると、、、
	fmt.Printf("func6=%d\n", func6())

	// 引数を無視する関数
	fmt.Printf("func7=%d\n", func7(2, 3))

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("f type=%T, f=%d\n", f, f(2, 3))

	fmt.Printf("f type=%#v\n", func(a, b int) int { return a + b }) // アドレスが返る。
	fmt.Printf("f value=%d\n", func(a, b int) int { return a + b }(2, 3))

	// 関数を返す関数を呼ぶ
	f3 := func8()
	f3()

	// そのまま呼ぶ
	func8()()

	// 無名関数を引数に渡して呼ぶ
	func9(func() { fmt.Printf("no name function\n") })

	// クロージャ
	fc := func10closure()
	fmt.Printf("こんにちは->%v\n", fc("こんにちは"))
	fmt.Printf("さようなら->%v\n", fc("さようなら"))
	fmt.Printf("おはよう->%v\n", fc("おはよう"))

	// 別の変数経由ならどうなる？
	fc2 := func10closure()
	fmt.Printf("いってきます->%v\n", fc2("いってきます"))
	fmt.Printf("おやすみなさい->%v\n", fc2("おやすみなさい"))
	fmt.Printf("いただきます->%v\n", fc2("いただきます"))
	fmt.Printf("ただいま->%v\n", fc("ただいま"))
	// こちらは、別に変数が保有される。

}

// 各引数に型指定
func func1(x int, y int) int {
	return x + y
}

// こんな書き方も可能
func func2(x, y int) int {
	return x * y
}

// いろんな型が混ざる
func func3(a, b float64, c, d int) float64 {
	tmp := a + float64(c)
	tmp2 := b / float64(d)
	return tmp + tmp2
}

// 複数の値を返す
func func4(a, b int) (int, int) {
	return a + b, a * b
}

// エラーと一緒に返す
func func5(a, b int) (int, interface{}) {
	tmp := a + b
	//	return tmp, "error happened"
	return tmp, nil
}

// 戻り値の短縮系
func func6() (a int) {
	return
}

// 引数を無視
func func7(_, _ int) int {
	a := 7
	a += 5
	return a
}

// 関数を返す関数
func func8() func() {
	return func() {
		fmt.Printf("this is a function\n")
	}
}

// 関数を引数にとる関数
func func9(f func()) {
	f()
}

func func10closure() func(string) string {
	// 一つ前に与えられた文字列を保存<-クロージャ変数
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}
