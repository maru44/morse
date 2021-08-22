package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	ch_rcv := make(chan string)
	ch_cancel := make(chan string)

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// close も入れる
	go func() {
		select {
		case v := <-ch_cancel:
			if v == "STOP" {
				break
			}
		default:
			for {
				select {
				case v := <-ch_rcv:
					fmt.Print("inp", v)
				case <-time.After(time.Second):
					fmt.Print("_")
				}
			}
		}
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		ch_rcv <- string(char)
		if key == keyboard.KeyEsc {
			break
		}
	}
}

func inputer(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)

	}()
	return ch
}

func oneWord3() string {
	// ch_inp := make(chan string)
	ch_inp := inputer(os.Stdin)

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
