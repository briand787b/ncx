package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.New())

	for {
		fmt.Println("hello world")
		time.Sleep(1000 * time.Millisecond)
	}
}
