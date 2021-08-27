package main

import (
	"morse/config"
	"morse/pkg/channel"
	"morse/pkg/file"
	"morse/pkg/mykey"

	"github.com/eiannone/keyboard"
)

func main() {
	ch := make(chan string)
	ret := ""

	if config.INPUT_MODE == "KEYBOARD" {
		mykey.InitKeyboard()
		defer keyboard.Close()
	}

	go channel.GeneralChanIn(&ret, ch)

	if config.INPUT_MODE == "KEYBOARD" {
		channel.OutKeyBoard(ch)
	}

	if config.OUTPUT_MODE == "TEXTFILE" {
		file.SaveFileFromConsole(ret)
	}
}
