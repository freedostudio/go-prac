package main

import (
	"fmt"

	"github.com/freedostudio/go-prac/dog"
)

func main() {
	fmt.Println("Hello, world.")
	aDog := dog.Dog{Kind: "heibei", Name: "xiaohei", Location: "beijing"}
	aDog.MakeSound()
	aDog.Run()
}
