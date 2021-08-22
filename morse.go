package main

import (
	"bufio"
	"context"
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

	// r := os.Stdin
	// fmt.Print(oneWord4(r))

	// for {
	// 	// inputer(r)
	// 	inp(r)
	// }

	// var a string
	aword := make([]byte, 1)
	fmt.Scan(&aword)
	fmt.Println(fmt.Printf("unko: %c", aword))
}

func inp(r io.Reader) {
}

func inputer(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			fmt.Print(s.Text() + "dddd")
			// ch <- s.Text()
		}
		close(ch)

	}()
	return ch
}

func getOneOrTimeOut() (str string) {
	r := os.Stdin
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		ch <- s.Text()
	}()

	go func() {
		select {
		case <-ch:
			str = "a"
			close(ch)
		}
	}()

	if str != "" {
		return
	}

	time.Sleep(time.Second)
	return ""
}

func interval(d time.Duration) <-chan string {
	ch := make(chan string)
	time.NewTimer(d)
	ch <- config.INTERVAL_LETTER
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

	go func() {
		select {
		case <-ch_inp:
			str = "a"
			return
			// case <-timer.C:
			// 	str = "b"
			// 	return
		}
	}()

	time.Sleep(time.Second)
	return "b"
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

func oneWord3() string {
	ch_inp := make(chan string)

	for {
		select {
		case <-ch_inp:
			// close(tl)
			fmt.Print("kita")
			return "a"
		case <-time.After(time.Second):
			return "b"
		}
	}
}

func oneWord4(r io.Reader) string {
	var ch_inp = inputer(r)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return "fin"
		case <-ch_inp:
			return "a"
		}
	}
}
