package morse

import (
	"fmt"
	"time"
)

func BaseReceive(m *Morse, ch chan string, ret *string, stdout bool) {
	for {
		select {
		case v := <-ch:
			if v == m.Quit {
				close(ch)
				break
			}
			m.ConvertInputCode(v, ret, stdout)
		case <-time.After(time.Duration(m.IntervalDuration) * time.Millisecond):
			*ret += m.Interval
			fmt.Print(m.Interval)
		}
	}
}
