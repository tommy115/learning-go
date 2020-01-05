package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName()) /* 関数AppName呼出 */

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
