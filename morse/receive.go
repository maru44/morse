package morse

import (
	"fmt"
	"time"
)

func BaseReceive(m *Morse, ch chan string, ret *string) {
	for {
		select {
		case v := <-ch:
			if v == m.QuitLetter {
				close(ch)
				break
			} else {
				m.ConvertInputCode(v, ret, true)
			}
		case <-time.After(time.Duration(m.Interval) * time.Millisecond):
			*ret += m.IntervalLetter
			fmt.Print(m.IntervalLetter)
		}
	}
}
