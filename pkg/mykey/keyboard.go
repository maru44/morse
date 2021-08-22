package mykey

import (
	"morse/config"
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
