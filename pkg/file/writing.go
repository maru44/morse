package file

import (
	"morse/config"
	"os"
	"strings"
)

func WriteFile(fileName, content string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteContent := []byte(content)

	file.Write(byteContent)
}

func ConvertCode(in string) string {
	var spaceLetter string
	if config.INTERVAL_MODE == "NORMAL" {
		spaceLetter = strings.Repeat(" ", 7)
	} else if config.INTERVAL_MODE == "SPEED" {
		spaceLetter = strings.Repeat(" ", 3)
	}
	ret := strings.Replace(in, spaceLetter, "/", -1)
	return ret
}
