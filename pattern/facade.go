package main

import "fmt"

/*
	Фасад предоставляет простой интерфейс к сложной системе которая состоит из множества разнообразных классов.
	Фасад применяется когда нужно уменьшить количество зависимостей между клиентом и сложной системой, когда нужно разделить систему на отдельные слои
	Плюс: Упрощает взаимодействие со сложной системой
	Минус: Риск стать god object

	В примере: реализован фасад для аудиосистемы. Аудио система состоит из компонентов: Цифро-Аналоговый преобразователь,
	предусилитель, усилитель и динамик. Для воспроизведения музыки нужно обработать цифровой сигнал на каждом из
	компонентов и подать аналоговый сигнал на динамик
*/

type Amp struct {
}

func (a *Amp) amp() {
	fmt.Println("Sound amped")
}

type PreAmp struct {
}

func (a *PreAmp) preAmp() {
	fmt.Println("Sound preamped")
}

type Speaker struct {
}

func (s *Speaker) getSound() {
	fmt.Println("*music is playing*")
}

type DCConverter struct {
}

func (d *DCConverter) convert() {
	fmt.Println("Signal converted to analog")
}

type audioSystemFacade struct {
	dc     *DCConverter
	preamp *PreAmp
	amp    *Amp

	speaker *Speaker
}

func (a *audioSystemFacade) playMusic(songName string) {
	fmt.Printf("Starting play music...\nSong: %s\n", songName)
	a.dc.convert()
	a.preamp.preAmp()
	a.amp.amp()
	a.speaker.getSound()
}

func NewAudioSystem() *audioSystemFacade {
	return &audioSystemFacade{
		dc:      &DCConverter{},
		preamp:  &PreAmp{},
		amp:     &Amp{},
		speaker: &Speaker{},
	}
}

func main() {
	as := NewAudioSystem()
	as.playMusic("")
}
