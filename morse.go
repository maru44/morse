package main

import (
	"fmt"
	"morse/pkg/keyboard"
)

func main() {
	// keyboard.InputEventOrIntervalFail()

	// go func() {
	// 	for {
	// 		fmt.Print(keyboard.InputEventOrInterval())
	// 	}
	// }()
	// time.Sleep(11 * time.Second)
	resultString := ""

	ret := keyboard.MyGetKey(resultString)
	fmt.Printf("result: %s", ret)
}
