package chapter3_decorator

import "fmt"

type Beverage interface {
	getDescription() string
	cost() float32
	GetSize() Size
}

type Size int

const (
	Tall Size = iota
	Grand
	Venti
)

func (s Size) String() string {
	switch s {
	case Tall:
		return "Tall"
	case Grand:
		return "Grand"
	case Venti:
		return "Venti"
	default:
		return " "
	}
}

type CommonBeverage struct {
	description string
	price       float32
	size        Size
}

func (c *CommonBeverage) getDescription() string {
	return fmt.Sprintf("%v %s", c.size, c.description)
}

func (c *CommonBeverage) cost() float32 {
	return c.price
}

func (c *CommonBeverage) SetSize(s Size) {
	c.size = s
}

func (c *CommonBeverage) GetSize() Size {
	return c.size
}

type DarkRoast struct {
	*CommonBeverage
}

func NewDarkRoast(s Size) Beverage {
	return &DarkRoast{&CommonBeverage{
		"DarkRoast", 1, s,
	}}
}

type Expresso struct {
	*CommonBeverage
}

func NewExpresso(s Size) Beverage {
	return &Expresso{&CommonBeverage{
		"Expresso", 2.0, s,
	}}
}
