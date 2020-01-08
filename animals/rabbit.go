package animals

import "fmt"

func init() {
	fmt.Println("init in animals.rabbit")
}

func RabbitFeed() string {
	return "Carrot"
}
