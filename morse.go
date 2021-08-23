package main

import (
	"bufio"
	"fmt"
	"morse/config"
	"morse/pkg/file"
	"morse/pkg/mykey"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	ch_rcv := make(chan string)
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
		for {
			select {
			case v := <-ch_rcv:
				if v == config.QUIT_LETTER {
					close(ch_rcv)
					break
				} else {
					res := mykey.ConvertInputCode(v)
					ret += res
					fmt.Print(res)
				}
			case <-time.After(config.TYPING_INTERVAL * time.Millisecond):
				ret += config.INTERVAL_LETTER
				fmt.Print(config.INTERVAL_LETTER)
			}
		}
	}()

	for {
		char, _, err := keyboard.GetKey()
		defer keyboard.Close()
		if err != nil {
			panic(err)
		}

		if string(char) == config.QUIT_PING {
			ch_rcv <- config.QUIT_LETTER
			keyboard.Close()
			break
		} else {
			ch_rcv <- string(char)
		}
	}

	// file save
	scan1 := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDo you save it? Press y OR n.")
	for {
		scan1.Scan()
		isSave := scan1.Text()

		switch isSave {
		case "y":
			fmt.Println("Enter file name.Default is morse.txt")
			scan2 := bufio.NewScanner(os.Stdin)
			scan2.Scan()
			fileName := scan2.Text()
			saveString := strings.TrimSpace(ret)
			if fileName == "" {
				file.WriteFile(
					fmt.Sprintf("%s%s.txt", config.DEFAULT_FILE_PATH, config.DEFAULT_FILE_NAME), saveString,
				)
			} else {
				file.WriteFile(
					fmt.Sprintf("%s%s.txt", config.DEFAULT_FILE_PATH, fileName), saveString,
				)
			}
			break
		default:
			break
		}
		break
	}
}
