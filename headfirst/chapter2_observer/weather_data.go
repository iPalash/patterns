package chapter2_observer

type WeatherData struct {
	Temperature int
	Pressure    int
	Humidity    int
	observers   []Observer
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) Changed(t, p, h int) {
	w.Temperature = t
	w.Pressure = p
	w.Humidity = h
	w.Notify()
}

func (w *WeatherData) AddObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, obs := range w.observers {
		if o == obs {
			w.observers = append(w.observers[:i-1], w.observers[i+1:]...)
			return
		}
	}
}

func (w *WeatherData) Notify() {
	for _, obs := range w.observers {
		obs.update()
	}
}
