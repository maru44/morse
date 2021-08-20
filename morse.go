package main

import (
	"bufio"
	"fmt"
	"io"
	"morse/config"
	"morse/pkg/keyboard"
	"os"
	"time"
)

func main() {

	var (
	// ch_inp = inputer(os.Stdin)
	// to     = time.After(20 * time.Second)
	)

	// result := ""

	// keyInput:
	// 	for {
	// 		over := time.After(config.TYPING_INTERVAL * time.Millisecond)
	// 		select {
	// 		case v := <-ch_inp:
	// 			str := keyboard.ConvertInputCode(v)
	// 			fmt.Print(str)
	// 			result += str
	// 			if str == config.QUIT_LETTER {
	// 				break keyInput
	// 			}
	// 		case <-over:
	// 			str := config.INTERVAL_LETTER
	// 			fmt.Print(str)
	// 			result += str
	// 		}
	// 	}
	// 	fmt.Println(result)
	fmt.Print(oneWord())
}

func inputer(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func oneSection() (str string) {
	var (
		ch_inp = inputer(os.Stdin)
		to     = time.After(time.Second)
	)
oneWord:
	for {
		select {
		case v := <-ch_inp:
			str = keyboard.ConvertInputCode(v)
			break oneWord
		case <-to:
			str = config.INTERVAL_LETTER
			break oneWord
		}
	}
	return
}

func oneWord() (str string) {
	ch_inp := inputer(os.Stdin)
	// strInput := make(chan string)

	go func() {
		select {
		// case <-strInput:
		case <-ch_inp:
			str = "a"
		}
	}()

	// ch_inp := inputer(os.Stdin)
	// strInput <- <-ch_inp

	if str != "" {
		fmt.Print("inputed")
		return
	}

	time.Sleep(time.Second)
	str = "b"
	return
}

func oneWord2() (str string) {
	var ch_inp = inputer(os.Stdin)
	to := make(chan bool)

	// var to = time.After(5 * time.Second)
	// word:
	// time.NewTimer(5 * time.Second)
	// to <- true

	timer := time.NewTimer(5 * time.Second)
	to <- true

	for {
		select {
		case <-ch_inp:
			str = "a"
			timer.Stop()
			return
		case <-to:
			str = "b"
			return
		}
	}
}
