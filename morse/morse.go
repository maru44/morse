package morse

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

type (
	Send    func(chan string)
	Recieve func(chan string, *string)

	Morse struct {
		// default is j
		DitPing string
		// default is k
		DahPing string
		// default is l
		QuitPing string

		// default is .
		Dit string
		// default is -
		Dah string
		// default is " "
		Interval string
		// default is "QUIT"
		Quit string

		// one time unit long (millisecond)
		// shortgap (between letters): three time units long
		// medium gap (between words): seven time units long
		IntervalDuration int

		// send to chan
		Send
		// receive chan
		Recieve
	}
)

func (m *Morse) InitMessage() string {
	return fmt.Sprintf(
		"%s => '%s', %s => '%s'\n%s => to quit\none time unit is %d milliseconds",
		m.DitPing, m.Dit, m.DahPing, m.Dah,
		m.QuitPing, m.IntervalDuration,
	)
}

func (m *Morse) ConvertInputCode(in string, ret *string, stdout bool) {
	switch in {
	case m.DitPing:
		if stdout {
			fmt.Print(m.Dit)
		}
		*ret += m.Dit
	case m.DahPing:
		if stdout {
			fmt.Print(m.Dah)
		}
		*ret += m.Dah
	}
}

func (m *Morse) ConvertCode(in string) []byte {
	space := strings.Repeat(" ", 7)
	spaceReplaced := strings.Replace(in, space, " ....... ", -1)

	h := morse.NewHacker()
	out, err := h.Decode(strings.NewReader(spaceReplaced))
	if err != nil {
		panic(err)
	}
	return out
}

func (m *Morse) SetRecieve(f func(m *Morse, ch chan string, ret *string)) {
	m.Recieve = Recieve(func(ch chan string, ret *string) {
		f(m, ch, ret)
	})
}

func (m *Morse) SetSend(f func(m *Morse, ch chan string)) {
	m.Send = Send(func(ch chan string) {
		f(m, ch)
	})
}
