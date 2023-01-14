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
		SinglePing string
		// default is k
		TriplePing string
		// default is l
		QuitPing string

		// default is .
		SingleLetter string
		// default is -
		TripleLetter string
		// default is " "
		IntervalLetter string
		// default is "QUIT"
		QuitLetter string

		// Milli second
		Interval int

		// send to chan
		Send
		// receive chan
		Recieve
	}
)

func (m *Morse) InitMessage() string {
	return fmt.Sprintf(
		"%s => '%s', %s => '%s'\n%s => to quit\ninterval %d millisecond",
		m.SinglePing, m.SingleLetter, m.TriplePing, m.TripleLetter,
		m.QuitPing, m.Interval,
	)
}

func (m *Morse) ConvertInputCode(in string, ret *string, stdout bool) {
	switch in {
	case m.SinglePing:
		if stdout {
			fmt.Print(m.SingleLetter)
		}
		*ret += m.SingleLetter
	case m.TriplePing:
		if stdout {
			fmt.Print(m.TripleLetter)
		}
		*ret += m.TripleLetter
	}
}

func (m *Morse) ConvertCode(in string) []byte {
	spaceLetter := strings.Repeat(" ", 7)
	spaceToSlash := strings.Replace(in, spaceLetter, "/", -1)

	h := morse.NewHacker()
	out, err := h.Decode(strings.NewReader(spaceToSlash))
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
