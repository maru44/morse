package domain

import "github.com/maru44/morse/config"

type Morse struct {
	Settings MorseSetting
}

type InputMode string

type OutputMode string

type InputInterval int

type MorseSetting struct {
	Input    InputMode
	Output   OutputMode
	Interval InputInterval
}

type MorseInteractor interface {
	Ignition()
	SendChan(targetChannel chan<- string)
	ReceiveChanWithConvert(targetPointer *string, ch chan string)
	ReturnLetters(target string) (morseString string, decoded string)
}

func InitMorse(settings ...MorseSetting) Morse {
	var s MorseSetting
	if settings == nil {
		// default settings
		s = MorseSetting{
			Input:    InputMode(config.INPUT_MODE_KEYBOARD),
			Output:   OutputMode(config.OUTPUT_MODE_TEXTFILE),
			Interval: config.TYPING_INTERVAL,
		}
	} else {
		s = settings[0]
		if s.Interval == 0 {
			s.Interval = config.TYPING_INTERVAL
		}
	}

	ret := Morse{
		Settings: s,
	}
	return ret
}
