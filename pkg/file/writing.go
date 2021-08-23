package file

import (
	"os"
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
