package main

import (
	"fmt"

	"./animals"
	"./sample"
)

func init() {
	fmt.Println("init in main")
}

func main() {

	const (
		First = iota
		Number
		Control
		Assersion
		Assersion2
		Variable
		Variable2
		Variable3
		Func1
		GoRoutine
	)

	choice := GoRoutine

	switch choice {
	case First:
		fmt.Println(AppName()) /* 関数AppName呼出 */

		fmt.Println(animals.ElephantFeed())
		fmt.Println(animals.MonkeyFeed())
		fmt.Println(animals.RabbitFeed())
	case Variable:
		sample.VariableSample()
	case Variable2:
		sample.ChangeVariable()
	case Variable3:
		sample.MathVariable()
	case Number:
		sample.NumberSample()
	case Func1:
		sample.FuncSample()

	case Control:
		sample.ControlSample()

	case Assersion:

		sample.Anything(1)
		sample.Anything(3.14)
		sample.Anything(4 + 5i)
		sample.Anything('海')
		sample.Anything("海") //区別される
		sample.Anything("日本語")
		sample.Anything2(3)

	case Assersion2:
		sample.AssersionSample()

	case GoRoutine:
		sample.ParallelSample()

	default:
		fmt.Println("unknown case")
	}
}
