package strategy1

import "testing"

func TestKing_fight(t *testing.T) {
	tests := []struct {
		name string
		f    func() Character
	}{
		{"King with sword", NewKing},
		{"Magician with wand", NewMagician},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.f()
			k.fight()
		})
	}
}
