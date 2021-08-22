package main

import (
	"fmt"
	"morse/config"
	"morse/pkg/mykey"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	ch_rcv := make(chan string)
	ch_cancel := make(chan string)
	ret := ""

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	fmt.Printf(
		"%s => '%s', %s => '%s'\n%s => to quit\ninterval %d millisecond\n",
		config.SINGLE_PING, config.SINGLE_LETTER, config.TRIPLE_PING, config.TRIPLE_LETTER,
		config.QUIT_PING, config.TYPING_INTERVAL,
	)
	defer keyboard.Close()

	// close も入れる
	go func() {
		select {
		case v := <-ch_cancel:
			if v == "STOP" {
				fmt.Println(ret)
				break
			}
		default:
			for {
				select {
				case v := <-ch_rcv:
					res := mykey.ConvertInputCode(v)
					ret += res
					fmt.Print(res)
				case <-time.After(config.TYPING_INTERVAL * time.Millisecond):
					ret += config.INTERVAL_LETTER
					fmt.Print(config.INTERVAL_LETTER)
				}
			}
		}
	}()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if string(char) == config.QUIT_PING {
			break
		} else {
			ch_rcv <- string(char)
		}
	}
}
