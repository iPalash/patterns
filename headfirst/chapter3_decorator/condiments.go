package chapter3_decorator

type Mocha struct {
	Beverage
}

func AddMocha(b Beverage) Beverage {
	return &Mocha{b}
}

func (m *Mocha) getDescription() string {
	return m.Beverage.getDescription() + `, Mocha`
}

func (m *Mocha) cost() float32 {
	cost := m.Beverage.cost()
	switch m.Beverage.GetSize() {
	case Tall:
		return cost + 0.1
	case Grand:
		return cost + 0.2
	case Venti:
		return cost + 0.3
	}
	return cost
}

type Milk struct {
	Beverage
}

func AddMilk(b Beverage) Beverage {
	return &Milk{b}
}

func (m *Milk) getDescription() string {
	return m.Beverage.getDescription() + `, Milk`
}

func (m *Milk) cost() float32 {
	return m.Beverage.cost() + 0.1
}

type Soy struct {
	Beverage
}

func AddSoy(b Beverage) Beverage {
	return &Soy{b}
}

func (m *Soy) getDescription() string {
	return m.Beverage.getDescription() + `, Soy`
}

func (m *Soy) cost() float32 {
	return m.Beverage.cost() + 0.05
}
