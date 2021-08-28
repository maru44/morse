package execute

import (
	"fmt"
	"morse/config"
	"morse/pkg/mykey"
	"time"
)

func (mr *MorseRepository) receiveChanWithPrint(targetP *string, ch chan string) {
	for {
		select {
		case v := <-ch:
			if v == config.QUIT_LETTER {
				close(ch)
				break
			} else {
				res := mykey.ConvertInputCode(v)
				*targetP += res
				fmt.Print(res)
			}
		case <-time.After(config.TYPING_INTERVAL * time.Millisecond):
			*targetP += config.INTERVAL_LETTER
			fmt.Print(config.INTERVAL_LETTER)
		}
	}
}
