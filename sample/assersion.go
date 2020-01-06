package sample

import (
	"fmt"
)

func AssersionSample() {
	var x interface{}
	Anything3(x)
	Anything3(10)
	Anything3("hello")
	Anything3(5 + 3i)

	//型でswitchさせる
	anything4(5)
	anything4(3.2)
	anything4("山")
	anything4('山')
	anything4(3 + 7i)
}

func Anything(a interface{}) {
	fmt.Println(a)
}

func Anything2(x interface{}) {
	i := x.(int)
	fmt.Printf("changed to int %v", i)
	f := x.(float64)
	fmt.Printf("changed to float64 %v", f)

}

func Anything3(x interface{}) {
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unknown type")
	}
}

// 型アサーション
func anything4(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Println("this is bool")
	case int:
		fmt.Println("this is int")
	case string:
		fmt.Println("this is string")
	default:
		fmt.Println("unknown")
	}
}
