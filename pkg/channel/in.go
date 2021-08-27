package channel

import (
	"fmt"
	"morse/config"
	"morse/pkg/mykey"
	"time"
)

func GeneralChanIn(ret *string, ch chan string) string {
	for {
		select {
		case v := <-ch:
			if v == config.QUIT_LETTER {
				close(ch)
				break
			} else {
				res := mykey.ConvertInputCode(v)
				*ret += res
				fmt.Print(res)
			}
		case <-time.After(config.TYPING_INTERVAL * time.Millisecond):
			*ret += config.INTERVAL_LETTER
			fmt.Print(config.INTERVAL_LETTER)
		}
	}
}
