package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Vitor"
	var version float32 = 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is at version", version)

	fmt.Println("The 'name' variable type is", reflect.TypeOf(name))
}
