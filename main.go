package main

import (
	"fmt"
	feature "study/feature1"
	"study/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("githw1 !!!")

	feature.Feature1("feature1 !!!")

	simple_connection.CheckConnection()
}
