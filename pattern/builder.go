package main

import (
	"fmt"
	"strconv"
)

/*
	Строитель - порождающий паттерн позволяющий создавать объекты с различными параметрами шаг за шагом
		используя один и тот же процесс строительства.
	Используется когда целевой объект сложен и для его создания используется множество этапов
	Плюсы: Для создания объектов используется один и тот же код. Изоляция от сложной сборки
	Минусы: Усложняет код добавлением новых классов
*/

type Builder interface {
	setBody()
	setEngine()
	setSpareWheel()
	getCar() *Car
}

func getCarBuilder(carType string) Builder {
	switch carType {
	case "passenger car":
		return newPassengerCarBuilder()
	case "suv":
		return newSUVCarBuilder()
	default:
		return nil
	}
}

////////////////////////////////////////////////////////////////
type PassCarBuilder struct {
	body       string
	engine     string
	spareWheel bool
}

func newPassengerCarBuilder() *PassCarBuilder {
	return &PassCarBuilder{}
}

func (p *PassCarBuilder) setBody() {
	p.body = "passenger body"
}
func (p *PassCarBuilder) setSpareWheel() {
	p.spareWheel = false
}
func (p *PassCarBuilder) setEngine() {
	p.engine = "1.5 L"
}
func (p *PassCarBuilder) getCar() *Car {
	return &Car{
		p.body,
		p.engine,
		p.spareWheel,
	}
}

////////////////////////////////////////////////////////////////

type SUVCarBuilder struct {
	body       string
	engine     string
	spareWheel bool
}

func newSUVCarBuilder() *SUVCarBuilder {
	return &SUVCarBuilder{}
}

func (s *SUVCarBuilder) setBody() {
	s.body = "suv body"
}
func (s *SUVCarBuilder) setSpareWheel() {
	s.spareWheel = true
}
func (s *SUVCarBuilder) setEngine() {
	s.engine = "3 L"
}

func (s *SUVCarBuilder) getCar() *Car {
	return &Car{
		s.body,
		s.engine,
		s.spareWheel,
	}
}

////////////////////////////////////////////////////////////////

type Car struct {
	body       string
	engine     string
	spareWheel bool
}

func main() {
	builder := getCarBuilder("suv")

	builder.setBody()
	builder.setEngine()
	builder.setSpareWheel()
	car := builder.getCar()
	fmt.Printf("SUV parameters:\nbody: %s\nengine: %s\nsparewheel: %s\n\n", car.body, car.engine, strconv.FormatBool(car.spareWheel))
	builder = getCarBuilder("passenger car")

	builder.setBody()
	builder.setEngine()
	builder.setSpareWheel()
	car = builder.getCar()
	fmt.Printf("Passenger car parameters:\nbody: %s\nengine: %s\nsparewheel: %s\n", car.body, car.engine, strconv.FormatBool(car.spareWheel))
}
