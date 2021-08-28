package domain

type Morse struct {
	settings MorseSetting
	MorseFunctions
}

type InputMode string

type OutputMode string

type MorseSetting struct {
	Input  InputMode
	Output OutputMode
}

type MorseFunctions interface{}

func InitMorse(settings ...MorseSetting) Morse {
	var s MorseSetting
	if settings == nil {
		// default settings
		s = MorseSetting{}
	} else {
		s = settings[0]
	}

	ret := Morse{
		settings: s,
	}
	return ret
}
