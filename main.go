package main

import (
	"fmt"
	"github.com/Numbnut68/244/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	milo := canine{
		name: "Milo",
		age:  dog.ConvertToDog(15),
	}
	fmt.Println(milo)
}
