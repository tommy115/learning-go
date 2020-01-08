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
