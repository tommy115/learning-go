package sample

import (
	"fmt"
	"log"
	"os"
)

func RefDataSample() {
	logger := log.New(os.Stdout, "[RefDataSample] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("slice sample")
	sliceSample()

	logger.Println("slice extend sample")
	sliceExtend()

	logger.Println("another slice sample")
	anotherSlice()

	logger.Println("str slice sample")
	strSlice()

	logger.Println("append slice")
	appendSlice()

	logger.Println("copySlice")
	copySclice()
}

func sliceSample() {
	s := make([]int, 3)
	fmt.Println(s)

	s[0] = 3
	fmt.Println(s)

	s[1] = 10
	fmt.Println(s)

	s[2] = 5
	fmt.Println(s)

	// これはpanicになる。
	// s[3] = 7

	fmt.Printf("len = %d\n", len(s))
}

func sliceExtend() {
	s := make([]float64, 5, 10)
	fmt.Printf("s=%#v\n", s)
	fmt.Printf("s len=%#v\n", len(s))
	fmt.Printf("s cap=%#v\n", cap(s))

	s[0] = 3.14
	fmt.Printf("s=%#v\n", s)

	/*
		s[6] = 2.8
		fmt.Printf("s=%#v\n", s)
	*/
}

func anotherSlice() {
	s := []int{1, 2, 3, 10, 100}
	fmt.Printf("s=%#v\n", s)

	// 配列だと,,,
	x := [...]int{1, 2, 5, -4, -3, -7}
	fmt.Printf("x=%#v\n", x)

	a := s[0:2]
	fmt.Printf("a=%#v\n", a)

	// 配列からスライスを作る
	y := x[:]
	fmt.Printf("y=%#v\n", y)
}

func strSlice() {
	c := "abcde"
	c2 := c[0:3]
	fmt.Printf("c2=%#v\n", c2)

	s := "あいうえお"
	s2 := s[0:1]
	fmt.Printf("s2=%#v\n", s2) // UTF-8(3byte)なのでぶ分断される

}

func appendSlice() {
	s := []int{1, 2, 3}

	// 1要素追加
	s = append(s, 3)
	fmt.Println(s)

	// 3要素追加
	s = append(s, 5, 6, 7)
	fmt.Println(s)

	s1 := []int{3, 4, 5}
	s2 := []int{6, 7, 8}

	// sliceにsliceを追加する
	s1 = append(s1, s2...)
	fmt.Println(s1)

	// 文字列の場合
	//
	// ...が必要なのだが、、、、なんだこれ、、、
	//
	var b []byte
	b = append(b, "あいうえお"...)
	fmt.Println(b)

	b = append(b, "かきくけこ"...)
	fmt.Println(b)

	// 要素はappendだと、勝手に増えていく
	// (A) 要素サイズも確保サイズもなし
	a := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(a), cap(a))

	// (B) 数字を追加
	a = append(a, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(a), cap(a))

	// (C) 配列を追加
	a = append(a, []int{1, 2, 3}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(a), cap(a))

	// (D) 再び数字を追加
	a = append(a, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(a), cap(a))

	// (E) 複数数字を追加
	a = append(a, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(a), cap(a))

	a = append(a, 5)
	fmt.Printf("(F) len=%d, cap=%d\n", len(a), cap(a))

	a = append(a, 6, 7, 8, 9)
	fmt.Printf("(G) len=%d, cap=%d\n", len(a), cap(a))

	a2 := make([]int, 1024, 1024)
	fmt.Printf("a2 len=%d, cap=%d\n", len(a2), cap(a2))

	a2 = append(a2, 5)
	fmt.Printf("a2 len=%d, cap=%d\n", len(a2), cap(a2))
}

func copySclice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}

	fmt.Printf("n len=%d, value=%#v\n", len(s1), s1)

	// 元の要素分だけコピーされる
	// というか、srcとdstがシェルとかとは逆なのね、、、、
	n := copy(s1, s2)
	fmt.Printf("n len=%d, value=%#v, copied=%d\n", len(s1), s1, n)

	// コピー元の方が要素が多い場合、コピーできる要素数だけコピーされる
	s3 := []int{6, 7, 8}
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("n len=%d, value=%#v\n", len(s3), s3)

	n = copy(s3, s4)
	fmt.Printf("n len=%d, value=%#v, copied=%d\n", len(s3), s3, n)
}
