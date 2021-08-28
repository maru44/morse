package channel

import (
	"github.com/eiannone/keyboard"
	"github.com/maru44/morse/config"
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
