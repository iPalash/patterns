package chapter3_decorator

type Beverage interface {
	getDescription() string
	cost() float32
}

type CommonBeverage struct {
	description string
	price       float32
}

func (c *CommonBeverage) getDescription() string {
	return c.description
}

func (c *CommonBeverage) cost() float32 {
	return c.price
}

type DarkRoast struct {
	*CommonBeverage
}

func NewDarkRoast() Beverage {
	return &DarkRoast{&CommonBeverage{
		"DarkRoast", 1,
	}}
}

type Expresso struct {
	*CommonBeverage
}

func NewExpresso() Beverage {
	return &Expresso{&CommonBeverage{
		"Expresso", 2.0,
	}}
}
