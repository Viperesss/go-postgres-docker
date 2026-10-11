package main

import (
	"fmt"
	"my/feature1"
	simpleconnection "my/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Чууваак")

	feature1.Feature1()

	simpleconnection.CheckConnection()
}

func product() int {
	fmt.Println("мяясо")
	return 60
}
