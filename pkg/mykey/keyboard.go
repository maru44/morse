package mykey

import (
	"morse/config"
)

func ConvertInputCode(inp string) (out string) {
	if inp == config.SINGLE_PING {
		out = config.SINGLE_LETTER
	} else if inp == config.TRIPLE_PING {
		out = config.TRIPLE_LETTER
	} else if inp == config.QUIT_PING {
		out = config.QUIT_LETTER
	} else if inp == config.INTERVAL_LETTER {
		out = config.INTERVAL_LETTER
	} else {
		out = ""
	}
	return
}
