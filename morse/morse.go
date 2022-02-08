package morse

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/alwindoss/morse"
	"github.com/eiannone/keyboard"
)

type (
	SpeedMode string

	Pinger interface {
		Ping() string
		Letter() string
	}

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
		// default is NORMAL
		Speed SpeedMode
	}
)

const (
	SpeedModeFast   = SpeedMode("SPEED")
	SpeedModeNormal = SpeedMode("NORMAL")
)

func (m *Morse) InitMessage() string {
	return fmt.Sprintf(
		"%s => '%s', %s => '%s'\n%s => to quit\ninterval %d millisecond",
		m.SinglePing, m.SingleLetter, m.TriplePing, m.TripleLetter,
		m.QuitPing, m.Interval,
	)
}

func (m *Morse) ConvertInputCode(in string, ret *string) {
	switch in {
	case m.SinglePing:
		fmt.Print(m.SingleLetter)
		*ret += m.SingleLetter
	case m.TriplePing:
		fmt.Print(m.TripleLetter)
		*ret += m.TripleLetter
	}
}

func (m *Morse) ConvertCode(in string) string {
	spaceLetter := strings.Repeat(" ", 7)
	if m.Speed == SpeedModeFast {
		spaceLetter = strings.Repeat(" ", 3)
	}
	spaceToSlash := strings.Replace(in, spaceLetter, "/", -1)

	h := morse.NewHacker()
	out, err := h.Decode(strings.NewReader(spaceToSlash))
	if err != nil {
		panic(err)
	}
	return *(*string)(unsafe.Pointer(&out))
}

func (m *Morse) Recieve(ch chan string, ret *string) {
	for {
		select {
		case v := <-ch:
			if v == m.QuitLetter {
				close(ch)
				break
			} else {
				m.ConvertInputCode(v, ret)
			}
		case <-time.After(time.Duration(m.Interval) * time.Millisecond):
			*ret += m.IntervalLetter
			fmt.Print(m.IntervalLetter)
		}
	}
}

func (m *Morse) Send(ch chan string) {
	for {
		char, _, err := keyboard.GetKey()
		defer keyboard.Close()
		if err != nil {
			panic(err)
		}

		if string(char) == m.QuitPing {
			ch <- m.QuitLetter
			keyboard.Close()
			break
		} else {
			ch <- string(char)
		}
	}
}
