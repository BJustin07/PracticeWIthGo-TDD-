package main

import (
	"TrainingGo/helloWorld"
	"TrainingGo/integers"
	"fmt"
)

func main() {
	helloMessage := helloWorld.Hello("Luna", "Bisaya")
	fmt.Println(helloMessage)
	addition := integers.Add(2, 2)
	fmt.Println(addition)
}
