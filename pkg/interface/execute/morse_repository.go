package execute

import (
	"morse/config"
	"morse/pkg/domain"
	"morse/pkg/file"
	"morse/pkg/mykey"
)

type MorseRepository struct {
	domain.Morse
}

func (mr *MorseRepository) Ignition() {
	switch mr.Morse.Settings.Input {
	case config.INPUT_MODE_KEYBOARD:
		mykey.InitKeyboard()
	case config.INPUT_MODE_BROWSER_BUTTON:
	default:
	}
}

func (mr *MorseRepository) SendChan(ch chan<- string) {
	switch mr.Morse.Settings.Input {
	case config.INPUT_MODE_KEYBOARD:
		sendKeyBoard(ch)
	}
}

func (mr *MorseRepository) ReceiveChanWithEdit(targetP *string, ch chan string) {
	switch mr.Morse.Settings.Output {
	default:
		mr.receiveChanWithPrint(targetP, ch)
	}
}

func (mr *MorseRepository) ReturnLetters(target string) (morseString string, decoded string) {
	switch mr.Morse.Settings.Output {
	case config.OUTPUT_MODE_TEXT_WITH_PRINT:
		return file.SaveFileFromConsole(target)
	default:
		return file.SaveFileFromConsole(target)
	}
}
