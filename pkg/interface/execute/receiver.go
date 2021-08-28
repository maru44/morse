package execute

import (
	"fmt"
	"time"

	"github.com/maru44/morse/config"
	"github.com/maru44/morse/pkg/mykey"
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
		case <-time.After(time.Duration(mr.Settings.Interval) * time.Millisecond):
			*targetP += config.INTERVAL_LETTER
			fmt.Print(config.INTERVAL_LETTER)
		}
	}
}
