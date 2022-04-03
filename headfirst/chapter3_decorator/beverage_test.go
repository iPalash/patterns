package chapter3_decorator

import (
	"reflect"
	"testing"
)

func TestNewCoffee(t *testing.T) {
	t.Run("Espresso with milk", func(t *testing.T) {
		c := AddMilk(NewExpresso(Tall))
		want := "Tall Expresso, Milk"
		if got := c.getDescription(); !reflect.DeepEqual(got, want) {
			t.Errorf("NewExpresso() = %v, want %v", got, want)
		}
		if got := c.cost(); got != 2.1 {
			t.Errorf("NewExpresso() = %v, want %v", got, 2.1)
		}

	})

	t.Run("Darkroast with mocha and soy", func(t *testing.T) {
		c := AddSoy(AddMocha(NewDarkRoast(Tall)))
		want := "Tall DarkRoast, Mocha, Soy"
		if got := c.getDescription(); !reflect.DeepEqual(got, want) {
			t.Errorf("NewExpresso() = %v, want %v", got, want)
		}
		if got := c.cost(); got != 1.15 {
			t.Errorf("NewExpresso() = %v, want %v", got, 1.15)
		}

	})

	t.Run("Venti Darkroast with mocha", func(t *testing.T) {
		c := AddMocha(NewDarkRoast(Venti))
		want := "Venti DarkRoast, Mocha"
		if got := c.getDescription(); !reflect.DeepEqual(got, want) {
			t.Errorf("Coffee = %v, want %v", got, want)
		}
		if got := c.cost(); got != 1.3 {
			t.Errorf("COffee() = %v, want %v", got, 1.3)
		}

	})
}
