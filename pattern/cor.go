package main

import "fmt"

/*
	Цепочка вызовов - поведенческий паттерн, который позволяет передавать запросы последовательно по цепочке обработчиков.
	Применяется когда программе нужно выполнять запросы несколькими, заранее неизвестными, обработчиками
*/

type signal struct {
	s       []byte
	digital bool
	loud    bool
}

type Transformation interface {
	exec(*signal)
	setNext(Transformation)
}

type amp struct {
	next Transformation
}

func (a *amp) setNext(transformation Transformation) {
	a.next = transformation
}

func (a *amp) exec(s *signal) {
	fmt.Println("Sound amped")
	a.next.exec(s)
}

type preamp struct {
	next Transformation
}

func (p *preamp) exec(s *signal) {
	fmt.Println("Sound preamped")
	p.next.exec(s)
}
func (p *preamp) setNext(transformation Transformation) {
	p.next = transformation
}

type speaker struct {
	next Transformation
}

func (s *speaker) exec(sig *signal) {
	fmt.Println("*music is playing*")
}
func (s *speaker) setNext(transformation Transformation) {
	s.next = transformation
}

type dcconverter struct {
	next Transformation
}

func (d *dcconverter) setNext(transformation Transformation) {
	d.next = transformation
}
func (d *dcconverter) exec(s *signal) {
	fmt.Println("Signal converted to analog")
	d.next.exec(s)
}

func main() {
	speaker := &speaker{}

	dc := &dcconverter{}
	dc.setNext(speaker)

	amp := &amp{}
	amp.setNext(dc)

	preamp := &preamp{}
	preamp.setNext(amp)

	signal := &signal{
		s:       make([]byte, 10),
		digital: true,
		loud:    true,
	}
	preamp.exec(signal)

}
