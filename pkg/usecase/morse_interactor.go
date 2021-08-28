package usecase

type MorseInteractor struct {
	repo MorseRepository
}

type MorseRepository interface {
	Ignition()
	SendChan(chan<- string)
	ConvertTarget(*string)
	ReceiveChan(chan string)
	ReturnLetters(string) string
}

func (mi *MorseInteractor) Ignition() {
	mi.repo.Ignition()
}

func (mi *MorseInteractor) SendChan(targetChannel chan<- string) {
	mi.repo.SendChan(targetChannel)
}

func (mi *MorseInteractor) ReceiveChanWithConvert(targetP *string, ch chan string) {
	mi.repo.ConvertTarget(targetP)
	mi.repo.ReceiveChan(ch)
}

func (mi *MorseInteractor) ReturnLetters(target string) string {
	return mi.repo.ReturnLetters(target)
}
