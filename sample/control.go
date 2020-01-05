package sample

import (
	"fmt"
)

func LoopSample() {
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

func SwitchSample(n int) {

	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}
}

func SwitchFallSample(s string) {
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
