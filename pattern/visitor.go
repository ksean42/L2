package main

import "fmt"

/*
	Посетитель - поведенческий паттерн, позволяющий добавлять новую функциональность к объектам не изменяя их.
	Применяется когда необходимо выполнить какую либо операцию над множеством элементов различных типов которые нельзя изменять по тем или иным причинам.
	Плюсы: объединение однотипных операций над объектами. Может накапливать состояние при обходе множества элементов
Минусы: затруднено добавление новых классов, поскольку нужно обновлять иерархию посетителя и его сыновей.
*/

type Visitor interface {
	visitForPC(PC)
	visitForServer(Server)
}

type Computer interface {
	getType() string
	accept(v Visitor)
}

type PC struct {
	cpuCores int
	ramGB    int
	hddGB    int
}

func (p *PC) getType() string {
	return "PC"
}

func (p *PC) accept(v Visitor) {
	v.visitForPC(*p)
}

type Server struct {
	cpuCores int
	ramGB    int
	hddGB    int
}

func (s *Server) getType() string {
	return "PC"
}

func (s *Server) accept(v Visitor) {
	v.visitForServer(*s)
}

type ConcreteVisitor struct {
	performance string
}

func (c *ConcreteVisitor) visitForPC(p PC) {
	if p.cpuCores > 3 && p.ramGB > 8 && p.hddGB > 2000 {
		c.performance = "fast"
	} else {
		c.performance = "slow"
	}
	fmt.Println("PC:", c.performance)
}

func (c *ConcreteVisitor) visitForServer(s Server) {
	if s.cpuCores > 3 && s.ramGB > 8 && s.hddGB > 2000 {
		c.performance = "fast"
	} else {
		c.performance = "slow"
	}
	fmt.Println("Server:", c.performance)
}

func main() {
	pc := &PC{
		2,
		8,
		500,
	}
	server := &Server{
		16,
		32,
		5000,
	}

	concVisitor := &ConcreteVisitor{}
	pc.accept(concVisitor)
	server.accept(concVisitor)

}
