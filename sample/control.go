package sample

import (
	"fmt"
)

func ControlSample() {
	loopSample()

	switchSample(3)
	switchSample(1)
	switchSample(5)

	switchFallSample("A")
	switchFallSample("B")
	switchFallSample("CDE")

	gotoSample()
	labelSample()
	labelLoopSample()
}

func loopSample() {
	i := 0
	for {
		i++
		if i == 100 {
			fmt.Println(i)
			break
		}
	}

	i = 0
	for i < 100 {
		i++
		fmt.Println(i)
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
}

func switchSample(n int) {

	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}
}

func switchFallSample(s string) {
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)
}

func gotoSample() {
	for {
		for {
			for {
				fmt.Println("inner loop")
				goto DONE
			}
		}
	}
DONE:
	fmt.Println("done")
}

func labelSample() {
LOOP:
	for {
		for {
			for {
				fmt.Println("開始")
				break LOOP
			}
			fmt.Println("通らない") // 通らない箇所は警告になる
		}
		fmt.Println("通らない") // 通らない箇所は警告になる
	}
	fmt.Println("完了")
}

func labelLoopSample() {
L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは実行されない") // 通らないけど、これは警告にならない、、、、
	}
}
