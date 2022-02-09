package morse

type Option func(*Morse)

func NewMorse(opts ...Option) *Morse {
	m := &Morse{
		SinglePing:                   "j",
		TriplePing:                   "k",
		QuitPing:                     "l",
		SingleLetter:                 ".",
		TripleLetter:                 "-",
		IntervalLetter:               " ",
		QuitLetter:                   "QUIT",
		Interval:                     400,
		Speed:                        SpeedModeNormal,
		DefaultSavingFileDir:         "./storage/",
		DefaultSavingFileName:        "morse.txt",
		DefaultSavingFileDecodedName: "morse_decode.txt",
	}

	for _, o := range opts {
		o(m)
	}

	return m
}

func SinglePing(s string) Option {
	return func(m *Morse) {
		m.SinglePing = s
	}
}

func TriplePing(s string) Option {
	return func(m *Morse) {
		m.TriplePing = s
	}
}

func QuitPing(s string) Option {
	return func(m *Morse) {
		m.QuitPing = s
	}
}

func SingleLetter(s string) Option {
	return func(m *Morse) {
		m.SingleLetter = s
	}
}

func TripleLetter(s string) Option {
	return func(m *Morse) {
		m.TripleLetter = s
	}
}

func IntervalLetter(s string) Option {
	return func(m *Morse) {
		m.IntervalLetter = s
	}
}

func QuitlLetter(s string) Option {
	return func(m *Morse) {
		m.QuitLetter = s
	}
}

func Interval(in int) Option {
	return func(m *Morse) {
		m.Interval = in
	}
}

func Speed(s SpeedMode) Option {
	return func(m *Morse) {
		m.Speed = s
	}
}

func DefaultSavingFileDir(s string) Option {
	return func(m *Morse) {
		m.DefaultSavingFileDir = s
	}
}

func DefaultSavingFileName(s string) Option {
	return func(m *Morse) {
		m.DefaultSavingFileName = s
	}
}

func DefaultSavingFileDecodedName(s string) Option {
	return func(m *Morse) {
		m.DefaultSavingFileDecodedName = s
	}
}
