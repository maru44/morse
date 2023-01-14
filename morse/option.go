package morse

type Option func(*Morse)

func NewMorse(opts ...Option) *Morse {
	m := &Morse{
		SinglePing:     "j",
		TriplePing:     "k",
		QuitPing:       "l",
		SingleLetter:   ".",
		TripleLetter:   "-",
		IntervalLetter: " ",
		QuitLetter:     "QUIT",
		Interval:       400,
	}

	for _, o := range opts {
		o(m)
	}

	return m
}

// SinglePing is function to combine `.` and input key.
func SinglePing(s string) Option {
	return func(m *Morse) {
		m.SinglePing = s
	}
}

// TriplePing is function to combine `-` and input key.
func TriplePing(s string) Option {
	return func(m *Morse) {
		m.TriplePing = s
	}
}

// QuitPing is function to inform finish of input.
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
