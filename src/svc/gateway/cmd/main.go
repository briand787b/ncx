package main

import (
	"fmt"
	"time"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.New())

	cslog.NewCSLogger(nil, nil)

	for {
		fmt.Println("hello world")
		time.Sleep(1000 * time.Millisecond)
	}
}
