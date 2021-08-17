package keyboard

import (
	"fmt"
	"math/rand"
	"morse/config"
	"time"

	"github.com/eiannone/keyboard"
)

func InputEventOrIntervalFail() {
	ticker := time.NewTicker(config.TYPING_INTERVAL * time.Millisecond)
	defer ticker.Stop()

	fin := make(chan bool)

	go func() {
		for {
			select {
			// case 入力があったら:
			case <-fin:
				return
			case <-ticker.C:
				print("_")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Print("1")
	time.Sleep(3 * time.Second)
	fin <- true
	fmt.Print("finish")
}

func InputEventOrInterval() string {
	str := "_"
	ping := make(chan bool)

	go func() {
		select {
		case <-ping:
			str = "a"
		}
	}()

	// 仮のping送信
	rand.Seed(time.Now().UnixNano())
	if a := rand.Intn(2); a == 1 {
		ping <- true
	}

	if str == "a" {
		return str
	}

	time.Sleep(config.TYPING_INTERVAL * time.Millisecond)
	return str
}

func MyGetKey(ret string) string {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {

		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if char == []rune(config.STOP_PING)[0] {
			fmt.Println()
			break
		} else {
			ret += string(char)
			fmt.Printf(string(char))
		}
	}
	return ret
}
