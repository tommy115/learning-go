package sample

import (
	"fmt"
)

func Anything(a interface{}) {
	fmt.Println(a)
}

func Anything2(x interface{}) {
	i := x.(int)
	fmt.Println("changed to int %v", i)
	f := x.(float64)
	fmt.Println("changed to float64 %v", f)

}
