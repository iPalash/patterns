package chapter2_observer

type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	Notify()
}

type Observer interface {
	update()
}
