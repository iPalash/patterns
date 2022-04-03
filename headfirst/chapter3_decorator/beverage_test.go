package chapter3_decorator

import (
	"reflect"
	"testing"
)

func TestNewCoffee(t *testing.T) {
	t.Run("Espresso with milk", func(t *testing.T) {
		c := AddMilk(NewExpresso())
		if got := c.getDescription(); !reflect.DeepEqual(got, "Expresso, Milk") {
			t.Errorf("NewExpresso() = %v, want %v", got, "Expresso, Milk")
		}
		if got := c.cost(); got != 2.1 {
			t.Errorf("NewExpresso() = %v, want %v", got, 2.1)
		}

	})

	t.Run("Darkroast with mocha and soy", func(t *testing.T) {
		c := AddSoy(AddMocha(NewDarkRoast()))
		if got := c.getDescription(); !reflect.DeepEqual(got, "DarkRoast, Mocha, Soy") {
			t.Errorf("NewExpresso() = %v, want %v", got, "")
		}
		if got := c.cost(); got != 1.15 {
			t.Errorf("NewExpresso() = %v, want %v", got, 1.15)
		}
	})
}
