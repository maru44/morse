package morse

type IMorse interface {
	Morse() Morse
	Send(ch chan string)
	Recieve(ch chan string, ret *string)
}
