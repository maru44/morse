package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/alwindoss/morse"
	"github.com/maru44/morse/config"
)

func writeFile(fileName, content string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent := []byte(content)

	file.Write(byteContent)
}

func convertCode(in string) string {
	var spaceLetter string
	if config.INTERVAL_MODE == "NORMAL" {
		spaceLetter = strings.Repeat(" ", 7)
	} else if config.INTERVAL_MODE == "SPEED" {
		spaceLetter = strings.Repeat(" ", 3)
	}
	spaceToSlash := strings.Replace(in, spaceLetter, "/", -1)
	return convertDecoding(spaceToSlash)
}

func convertDecoding(in string) string {
	h := morse.NewHacker()
	out, err := h.Decode(strings.NewReader(in))
	if err != nil {
		panic(err)
	}
	return *(*string)(unsafe.Pointer(&out))
}

func SaveFileFromConsole(ret string) (morseString string, decoded string) {
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
			morseString = strings.TrimSpace(ret)
			decoded = convertCode(morseString)
			if fileName == "" {
				writeFile(
					fmt.Sprintf("%s%s.txt", config.DEFAULT_FILE_PATH, config.DEFAULT_FILE_NAME), morseString,
				)
				writeFile(
					fmt.Sprintf("%s%s_decode.txt", config.DEFAULT_FILE_PATH, config.DEFAULT_FILE_NAME), decoded,
				)
			} else {
				writeFile(
					fmt.Sprintf("%s%s.txt", config.DEFAULT_FILE_PATH, fileName), morseString,
				)
				writeFile(
					fmt.Sprintf("%s%s_decode.txt", config.DEFAULT_FILE_PATH, fileName), decoded,
				)
			}
			break
		default:
			break
		}
		break
	}
	return morseString, decoded
}
