package main

import (
	"fmt"
	"morse/pkg/keyboard"
	"time"
)

func main() {
	// keyboard.InputEventOrIntervalFail()

	go func() {
		for {
			fmt.Print(keyboard.InputEventOrInterval())
		}
	}()
	time.Sleep(11 * time.Second)
}
