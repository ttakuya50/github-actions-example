package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func main() {
	fmt.Println(sum(1, 3))
	fmt.Println(ToSnake())
}

func sum(a, b int) int {
	return a + b
}

func ToSnake() string {
	return strcase.ToSnake("UserStart")
}
