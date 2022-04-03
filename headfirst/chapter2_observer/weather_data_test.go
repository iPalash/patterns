package chapter2_observer

import "testing"

func TestWeatherData_Changed(t *testing.T) {
	wd := NewWeatherData()

	_ = NewTempDiplay(wd)
	_ = NewHumidityDisplay(wd)

	type args struct {
		t int
		p int
		h int
	}
	tests := []struct {
		name string
		args args
	}{
		{"simple", args{
			1, 2, 3,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd.Changed(tt.args.t, tt.args.p, tt.args.h)
		})
	}
}
