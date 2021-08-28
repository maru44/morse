package main

import (
	"morse/pkg/domain"
	"morse/pkg/interface/execute"
	"morse/pkg/mykey"
	"morse/pkg/usecase"
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
