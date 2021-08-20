package keyboard

import (
	"morse/config"
)

//
// func MyGetKey(ret string) string {
// 	if err := keyboard.Open(); err != nil {
// 		panic(err)
// 	}
// 	defer keyboard.Close()

// 	fmt.Println("Press L to quit")
// 	for {

// 		str := InputOrInterval()

// 		if str == config.QUIT_LETTER {
// 			fmt.Println()
// 			break
// 		} else {
// 			ret += str
// 			fmt.Print(str)
// 		}
// 	}
// 	return ret
// }

func ConvertInputCode(inp string) (out string) {
	if inp == config.SINGLE_PING {
		out = config.SINGLE_LETTER
	} else if inp == config.TRIPLE_PING {
		out = config.TRIPLE_LETTER
	} else if inp == config.QUIT_PING {
		out = config.QUIT_LETTER
	} else if inp == config.INTERVAL_LETTER {
		out = config.INTERVAL_LETTER
	} else {
		out = ""
	}
	return
}

// memo
/*
inputCommはgoroutine内じゃ受け取れない
GetKey()を使うとtime.Sleepもtime.Timerも効かない
*/
