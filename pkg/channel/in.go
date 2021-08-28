package channel

import (
	"fmt"
	"time"

	"github.com/maru44/morse/config"
	"github.com/maru44/morse/pkg/mykey"
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
