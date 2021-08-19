package keyboard

import (
	"errors"
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
	ping := make(chan string)

	char, _, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	go func() {
		t := time.NewTimer(config.TYPING_INTERVAL * time.Millisecond)
		<-t.C
		ping <- config.INTERVAL_LETTER
	}()

	go func() {
		for {
			select {
			case inp := <-ping:
				if inp == config.SINGLE_PING {
					str = config.SINGLE_LETTER
				} else if inp == config.TRIPLE_PING {
					str = config.TRIPLE_LETTER
				} else if inp == config.QUIT_PING {
					str = config.QUIT_LETTER
				} else if inp == config.INTERVAL_LETTER {
					fmt.Print("a")
					str = config.INTERVAL_LETTER
				}

			}
		}
	}()

	ping <- string(char)

	return
}

//
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

func getKey() (str string, err error) {
	waitingForKey := make(chan bool)
	inputComm := make(chan keyboard.KeyEvent)
	// Check if opened
	if !keyboard.IsStarted(time.Millisecond * 50) {
		return "", errors.New("keyboard not opened")
	}
	// Check if already waiting for key
	select {
	case waitingForKey <- true:
		return "", errors.New("already waiting for key")
	default:
	}

	// 動いてない

	for {
		go func() {
			select {
			case ev := <-inputComm:
				str = convertInputCode(string(ev.Rune))
				print(str)
				err = ev.Err

			case keepAlive := <-waitingForKey:
				if !keepAlive {
					str = ""
					err = errors.New("operation canceled")
					print(err)
				}
			}
		}()

		if str != "" {
			return
		}

		// こっちにしか来てない
		time.Sleep(config.TYPING_INTERVAL * time.Millisecond)
		str = config.INTERVAL_LETTER
		return
	}
}

func OriginalGetKey() (str string) {
	rawStr, _ := getKey()
	if rawStr == config.SINGLE_PING {
		str = config.SINGLE_LETTER
	} else if rawStr == config.TRIPLE_PING {
		str = config.TRIPLE_LETTER
	} else if rawStr == config.QUIT_PING {
		str = config.QUIT_LETTER
	} else if rawStr == config.INTERVAL_LETTER {
		str = config.INTERVAL_LETTER
	}
	return
}

func convertInputCode(inp string) (out string) {
	if inp == config.SINGLE_PING {
		out = config.SINGLE_LETTER
	} else if inp == config.TRIPLE_PING {
		out = config.TRIPLE_LETTER
	} else if inp == config.QUIT_PING {
		out = config.QUIT_LETTER
	} else if inp == config.INTERVAL_LETTER {
		out = config.INTERVAL_LETTER
	}
	return
}

// memo
/*
inputCommはgoroutine内じゃ受け取れない
GetKey()を使うとtime.Sleepもtime.Timerも効かない
*/
