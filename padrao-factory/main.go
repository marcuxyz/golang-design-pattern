package main

import (
	"fmt"
	"myapp/factory"
)

func main() {
	f := factory.Factory("senior")
	fmt.Println(f.Salario())
}
