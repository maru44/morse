package mykey

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/maru44/morse/config"
)

func ConvertInputCode(inp string) (out string) {
	if inp == config.SINGLE_PING {
		out = config.SINGLE_LETTER
	} else if inp == config.TRIPLE_PING {
		out = config.TRIPLE_LETTER
	} else {
		out = ""
	}
	return
}

func InitKeyboard() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	fmt.Printf(
		"%s => '%s', %s => '%s'\n%s => to quit\ninterval %d millisecond\n",
		config.SINGLE_PING, config.SINGLE_LETTER, config.TRIPLE_PING, config.TRIPLE_LETTER,
		config.QUIT_PING, config.TYPING_INTERVAL,
	)
}

func CloseKeyboard() {
	keyboard.Close()
}

func GetKey() (rune, keyboard.Key, error) {
	return keyboard.GetKey()
}
