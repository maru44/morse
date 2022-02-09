package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/maru44/morse/morse"
)

func main() {
	ch := make(chan string)
	r := ""
	ret := &r

	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	m := morse.NewMorse()
	m.SetSend(func(m *morse.Morse, ch chan string) {
		send(m, ch)
	})
	m.SetRecieve(func(m *morse.Morse, ch chan string, ret *string) {
		receive(m, ch, ret)
	})

	fmt.Println(m.InitMessage())

	go m.Recieve(ch, ret)
	m.Send(ch)

	// file save
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDo you save it? Press y OR n.")
	scan.Scan()
	isSave := scan.Text() == "y"

	// saving
	if isSave {
		fmt.Println("Enter file name. Default is morse.txt")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		fileName := scan.Text()
		saveString := strings.TrimSpace(*ret)
		if fileName == "" {
			if err := writeFile(m.DefaultSavingFileDir+m.DefaultSavingFileName, []byte(saveString)); err != nil {
				panic(err)
			}
			if err := writeFile(m.DefaultSavingFileDir+m.DefaultSavingFileDecodedName, m.ConvertCode(saveString)); err != nil {
				panic(err)
			}
			return
		}

		if err := writeFile(fmt.Sprintf("%s.txt", fileName), []byte(saveString)); err != nil {
			panic(err)
		}
		if err := writeFile(fmt.Sprintf("%s_decode.txt", fileName), m.ConvertCode(saveString)); err != nil {
			panic(err)
		}
	}
}

func send(m *morse.Morse, ch chan string) {
	for {
		char, _, err := keyboard.GetKey()
		defer keyboard.Close()
		if err != nil {
			panic(err)
		}

		if string(char) == m.QuitPing {
			ch <- m.QuitLetter
			keyboard.Close()
			break
		} else {
			ch <- string(char)
		}
	}
}

func receive(m *morse.Morse, ch chan string, ret *string) {
	for {
		select {
		case v := <-ch:
			if v == m.QuitLetter {
				close(ch)
				break
			} else {
				m.ConvertInputCode(v, ret)
			}
		case <-time.After(time.Duration(m.Interval) * time.Millisecond):
			*ret += m.IntervalLetter
			fmt.Print(m.IntervalLetter)
		}
	}
}

func writeFile(fileName string, content []byte) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.Write(content); err != nil {
		return err
	}
	return nil
}
