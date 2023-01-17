# Morse

Type Morse code!

## Demo

![demo](https://raw.github.com/wiki/maru44/morse/images/morse1.gif)

## How to start inputing by cli

`go run main.go`

## Settings

You can edit settings.

```go
package main

import "github.com/maru44/morse/morse"

func main() {
    // set with args
    m1 := morse.NewMorse(morse.DitPing("a"), morse.DahPing("b"))

    // overwrite
    m2 := morse.NewMorse()
    m2.DitPing = "a"
    m2.DahPing = "b"

    // define by struct
    m3 := morse.Morse{
        DitPing: "a",
        DahPing: "b",
        // ...
    }
}
```

The default input settings is the below.

```
"j" > . (dit)
"k" > - (dah)   (You don't have to press it long.)
"l" > to quit it

0.4sec > one time unit
```

## You can use morse as package.

Core system is in here (https://github.com/maru44/morse/morse).
You can use it as package.

`go install github.com/maru44/morse@latest`

## Thanks

https://github.com/eiannone/keyboard

https://github.com/alwindoss/morse
