package morse

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

type (
	SpeedMode string

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
		// default is NORMAL
		Speed SpeedMode

		// default ./storage/
		DefaultSavingFileDir string
		// default morse.txt
		DefaultSavingFileName string
		// default morse_decode.txt
		DefaultSavingFileDecodedName string

		// send to chan
		Send
		// receive chan
		Recieve
	}
)

const (
	SpeedModeFast   = SpeedMode("FAST")
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

func (m *Morse) ConvertCode(in string) []byte {
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
