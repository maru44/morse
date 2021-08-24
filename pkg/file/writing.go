package file

import (
	"morse/config"
	"os"
	"strings"
	"unsafe"

	"github.com/alwindoss/morse"
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
