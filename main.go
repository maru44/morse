package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/maru44/morse/config"
	"github.com/maru44/morse/morse"
)

type ms struct{}

// func (m *ms) Morse() *morse.Morse {
// 	return morse.NewMorse()
// }

// func (m *ms) Send() *morse.Morse {

// }

func main() {
	ch := make(chan string)
	r := ""
	ret := &r

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	m := morse.NewMorse()

	fmt.Println(m.InitMessage())
	defer keyboard.Close()

	go m.Recieve(ch, ret)
	m.Send(ch)

	// file save
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDo you save it? Press y OR n.")
	scan.Scan()
	isSave := scan.Text() == "y"

	if isSave {
		fmt.Println("Enter file name. Default is morse.txt")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		fileName := scan.Text()
		saveString := strings.TrimSpace(*ret)
		if fileName == "" {
			morse.WriteFile("storage/morse.txt", saveString)
			morse.WriteFile("storage/morse_decode.txt", m.ConvertCode(saveString))
			return
		}

		morse.WriteFile(
			fmt.Sprintf("%s%s.txt", config.DEFAULT_FILE_PATH, fileName), saveString,
		)
		morse.WriteFile(
			fmt.Sprintf("%s%s_decode.txt", config.DEFAULT_FILE_PATH, fileName), m.ConvertCode(saveString),
		)

	}
}
