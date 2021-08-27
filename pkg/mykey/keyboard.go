package mykey

import (
	"fmt"
	"morse/config"

	"github.com/eiannone/keyboard"
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
