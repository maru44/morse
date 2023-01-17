package morse

type Option func(*Morse)

func NewMorse(opts ...Option) *Morse {
	m := &Morse{
		DitPing:          "j",
		DahPing:          "k",
		QuitPing:         "l",
		Dit:              ".",
		Dah:              "-",
		Interval:         " ",
		Quit:             "QUIT",
		IntervalDuration: 400,
	}

	for _, o := range opts {
		o(m)
	}

	return m
}

// DitPing is a function to define key for `dit`.
// Default is `j`.
func DitPing(s string) Option {
	return func(m *Morse) {
		m.DitPing = s
	}
}

// DahPing is a function to define key for `dah`.
// Default is `k`.
func DahPing(s string) Option {
	return func(m *Morse) {
		m.DahPing = s
	}
}

// QuitPing is a function to define informing finish of input.
// Default is `l`.
func QuitPing(s string) Option {
	return func(m *Morse) {
		m.QuitPing = s
	}
}

// Dit is a function to set `dit`.
// Default is `.`.
func Dit(s string) Option {
	return func(m *Morse) {
		m.Dit = s
	}
}

// Dah is a function to set `dah`.
// Default is `-`.
func Dah(s string) Option {
	return func(m *Morse) {
		m.Dah = s
	}
}

// Interval is a function to set interval.
// Default is ` `.
func Interval(s string) Option {
	return func(m *Morse) {
		m.Interval = s
	}
}

// Quit is a function to define signal for finishing.
// Default is `QUIT`.
func Quit(s string) Option {
	return func(m *Morse) {
		m.Quit = s
	}
}

// IntervalDuration is a function to define one time unit long.
// Default is 400 millisecond.
func IntervalDuration(in int) Option {
	return func(m *Morse) {
		m.IntervalDuration = in
	}
}
