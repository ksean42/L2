package main

import (
	"fmt"
	"os"
)

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.
	В Go невозможно реализовать классический вариант паттерна Фабричный метод, поскольу в языке отсутствуют возможности ООП,
	в том числе классы и наследственность. Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.
*/

type Phone interface {
	setDisplay(string)
	setCpu(string)
	setManufacturer(string)
	printInfo()
}

type phone struct {
	display      string
	cpu          string
	manufacturer string
}

func (p *phone) setDisplay(displayType string) {
	p.display = displayType
}
func (p *phone) setCpu(cpuType string) {
	p.cpu = cpuType
}
func (p *phone) setManufacturer(brand string) {
	p.manufacturer = brand
}
func (p *phone) printInfo() {
	fmt.Printf("Phone:\nDisplay: %s\nCPU: %s\nBrand: %s\n", p.display, p.cpu, p.manufacturer)
}

type iPhoneX struct {
	phone
}

func newIPhone() Phone {
	return &iPhoneX{
		phone{"OLED",
			"Apple a11",
			"Apple",
		},
	}
}

type MiMix struct {
	phone
}

func newMiMIx() Phone {
	return &MiMix{
		phone{"IPS",
			"Snapdragon",
			"Xiaomi",
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]
	var p Phone
	if arg == "MiMix" {
		p = newMiMIx()
	} else if arg == "iPhone" {
		p = newIPhone()
	}
	if p != nil {
		p.printInfo()
	}
}
