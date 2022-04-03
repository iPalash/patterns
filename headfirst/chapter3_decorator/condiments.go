package chapter3_decorator

type Mocha struct {
	b Beverage
}

func AddMocha(b Beverage) Beverage {
	return &Mocha{b}
}

func (m *Mocha) getDescription() string {
	return m.b.getDescription() + `, Mocha`
}

func (m *Mocha) cost() float32 {
	return m.b.cost() + 0.1
}

type Milk struct {
	b Beverage
}

func AddMilk(b Beverage) Beverage {
	return &Milk{b}
}

func (m *Milk) getDescription() string {
	return m.b.getDescription() + `, Milk`
}

func (m *Milk) cost() float32 {
	return m.b.cost() + 0.1
}

type Soy struct {
	b Beverage
}

func AddSoy(b Beverage) Beverage {
	return &Soy{b}
}

func (m *Soy) getDescription() string {
	return m.b.getDescription() + `, Soy`
}

func (m *Soy) cost() float32 {
	return m.b.cost() + 0.05
}
