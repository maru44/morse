package usecase

import "github.com/maru44/morse/pkg/domain"

type MorseInteractor struct {
	repo MorseRepository
}

func NewMorseInteractor(mr MorseRepository) domain.MorseInteractor {
	return &MorseInteractor{
		repo: mr,
	}
}

type MorseRepository interface {
	Ignition()
	SendChan(chan<- string)
	ReceiveChanWithEdit(*string, chan string)
	ReturnLetters(string) (string, string)
}

func (mi *MorseInteractor) Ignition() {
	mi.repo.Ignition()
}

func (mi *MorseInteractor) SendChan(targetChannel chan<- string) {
	mi.repo.SendChan(targetChannel)
}

func (mi *MorseInteractor) ReceiveChanWithConvert(targetP *string, ch chan string) {
	mi.repo.ReceiveChanWithEdit(targetP, ch)
}

func (mi *MorseInteractor) ReturnLetters(target string) (string, string) {
	return mi.repo.ReturnLetters(target)
}
