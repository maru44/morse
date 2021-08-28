package execute

import (
	"morse/config"
	"morse/pkg/mykey"
)

func sendKeyBoard(c chan<- string) {
	for {
		char, _, err := mykey.GetKey()
		defer mykey.CloseKeyboard()
		if err != nil {
			panic(err)
		}

		if string(char) == config.QUIT_PING {
			c <- config.QUIT_LETTER
			mykey.CloseKeyboard()
			break
		} else {
			c <- string(char)
		}
	}
}
