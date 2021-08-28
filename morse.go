package main

import (
	"github.com/maru44/morse/pkg/domain"
	"github.com/maru44/morse/pkg/interface/execute"
	"github.com/maru44/morse/pkg/mykey"
	"github.com/maru44/morse/pkg/usecase"
)

func main() {
	ch := make(chan string)
	ret := ""

	morse := domain.InitMorse()
	mi := usecase.NewMorseInteractor(
		&execute.MorseRepository{Morse: morse},
	)

	mi.Ignition()
	defer mykey.CloseKeyboard()
	go mi.ReceiveChanWithConvert(&ret, ch)
	mi.SendChan(ch)
	mi.ReturnLetters(ret)
}
