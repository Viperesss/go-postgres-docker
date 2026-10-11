package main

import (
	"fmt"
	"my/feature1"
)

func main() {
	fmt.Println("Чууваак")

	feature1.Feature1()
}

func product() int {
	fmt.Println("мяясо")
	return 60
}
