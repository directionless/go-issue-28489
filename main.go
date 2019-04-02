package main

import (
	"fmt"

	"github.com/kolide/launcher/pkg/backoff"
)

func main() {
	bkoff := backoff.New()
	if err := bkoff.Run(sayHello); err != nil {
		panic(err)
	}
	fmt.Println("done")

}

func sayHello() error {
	fmt.Println("hello")
	return nil
}
