package main

import (
	"fmt"

	"./animals"
	"./sample"
)

func main() {

	const (
		First = iota
		Number
		Loop
		Switch
		Switch2
		Assersion
		Assersion2
		Variable
	)

	choice := Variable

	switch choice {
	case First:
		fmt.Println(AppName()) /* 関数AppName呼出 */

		fmt.Println(animals.ElephantFeed())
		fmt.Println(animals.MonkeyFeed())
		fmt.Println(animals.RabbitFeed())
	case Variable:
		sample.VariableSample()

	case Number:
		sample.NumberSample()

	case Loop:
		sample.LoopSample()

	case Switch:
		sample.SwitchSample(3)
		sample.SwitchSample(1)
		sample.SwitchSample(5)

	case Switch2:
		sample.SwitchFallSample("A")
		sample.SwitchFallSample("B")
		sample.SwitchFallSample("CDE")

	case Assersion:

		sample.Anything(1)
		sample.Anything(3.14)
		sample.Anything(4 + 5i)
		sample.Anything('海')
		sample.Anything("海") //区別される
		sample.Anything("日本語")

	case Assersion2:
		sample.Anything2(3)
	default:
		fmt.Println("unknown case")
	}
}
