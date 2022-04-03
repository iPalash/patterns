package chapter2_observer

import "fmt"

type Display interface {
	display()
}

type TempDisplay struct {
	wd *WeatherData
}

func NewTempDiplay(wd *WeatherData) Display {
	td := &TempDisplay{wd}
	wd.AddObserver(td)
	return td
}

func (d *TempDisplay) update() {
	d.display()
}

func (d *TempDisplay) display() {
	fmt.Println("Temp is:", d.wd.Temperature)
}

type HumidityDisplay struct {
	wd *WeatherData
}

func NewHumidityDisplay(wd *WeatherData) Display {
	hd := &HumidityDisplay{wd}
	wd.AddObserver(hd)
	return hd
}

func (d *HumidityDisplay) update() {
	d.display()
}

func (d *HumidityDisplay) display() {
	fmt.Println("Humidity is:", d.wd.Humidity)
}
