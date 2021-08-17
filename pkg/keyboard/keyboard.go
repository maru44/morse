package keyboard

import (
	"fmt"
	"math/rand"
	"morse/config"
	"time"
)

func InputEventOrIntervalFail() {
	ticker := time.NewTicker(config.TYPING_INTERVAL)
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

	rand.Seed(time.Now().UnixNano())
	if a := rand.Intn(2); a == 1 {
		ping <- true
	}

	time.Sleep(1 * time.Second)
	return str
}
