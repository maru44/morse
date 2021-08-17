package keyboard

import (
	"fmt"
	"math/rand"
	"morse/config"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	signal_single = "SINGLE"
	signal_triple = "TRIPLE"
	signal_quit   = "QUIT"
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

func InputOrInterval() (str string) {
	// ping := make(chan string)
	// go func() {
	// 	select {
	// 	case <-ping:
	// 		str = <-ping
	// 	}
	// }()

	char, _, err := keyboard.GetKey() // これでgoroutineが止まっている
	if err != nil {
		panic(err)
	}

	go func() {}()

	inp := string(char)
	if inp == config.SINGLE_PING {
		// ping <- config.SINGLE_LETTER
		// return
		return config.SINGLE_LETTER
	} else if inp == config.TRIPLE_PING {
		// ping <- config.TRIPLE_LETTER
		// return
		return config.TRIPLE_LETTER
	} else if inp == config.QUIT_PING {
		// ping <- config.QUIT_LETTER
		// return
		return config.QUIT_LETTER
	}

	time.Sleep(config.TYPING_INTERVAL * time.Millisecond)
	return config.INTERVAL_LETTER
}

func MyGetKey(ret string) string {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press L to quit")
	for {

		str := InputOrInterval()

		if str == config.QUIT_LETTER {
			fmt.Println()
			break
		} else {
			ret += str
			fmt.Print(str)
		}
	}
	return ret
}
