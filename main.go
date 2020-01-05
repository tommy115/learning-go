package main

import (
	"fmt"

	"./animals"
	"./sample"
)

func main() {

	const (
		SAMPLE_1ST = iota
		SAMPLE_LOOP
		SAMPLE_SWITCH
		SAMPLE_SWITCH2
		SAMPLE_ASSERSION
		SAMPLE_ASSERSION2
	)

	choice := SAMPLE_ASSERSION

	switch choice {
	case SAMPLE_1ST:
		fmt.Println(AppName()) /* 関数AppName呼出 */

		fmt.Println(animals.ElephantFeed())
		fmt.Println(animals.MonkeyFeed())
		fmt.Println(animals.RabbitFeed())

	case SAMPLE_LOOP:
		sample.LoopSample()

	case SAMPLE_SWITCH:
		sample.SwitchSample(3)
		sample.SwitchSample(1)
		sample.SwitchSample(5)

	case SAMPLE_SWITCH2:
		sample.SwitchFallSample("A")
		sample.SwitchFallSample("B")
		sample.SwitchFallSample("CDE")

	case SAMPLE_ASSERSION:

		sample.Anything(1)
		sample.Anything(3.14)
		sample.Anything(4 + 5i)
		sample.Anything('海')
		sample.Anything("海") //区別される
		sample.Anything("日本語")

	case SAMPLE_ASSERSION2:
		sample.Anything2(3)
	default:
		fmt.Println("unknown case")
	}
}
