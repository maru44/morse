package channel

import (
	"morse/config"

	"github.com/eiannone/keyboard"
)

func OutKeyBoard(c chan<- string) {
	for {
		char, _, err := keyboard.GetKey()
		defer keyboard.Close()
		if err != nil {
			panic(err)
		}

		if string(char) == config.QUIT_PING {
			c <- config.QUIT_LETTER
			keyboard.Close()
			break
		} else {
			c <- string(char)
		}
	}
}
