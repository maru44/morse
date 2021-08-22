package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var (
		ch_rcv = inputer(os.Stdin)
	)

	for {
		for {
			select {
			case v := <-ch_rcv:
				fmt.Print("inp", v)
			case <-time.After(time.Second):
				fmt.Print("_")
			}
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
