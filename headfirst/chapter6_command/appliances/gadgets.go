package appliances

import "fmt"

type TV struct {
}

func NewTV() *TV {
	return &TV{}
}

func (t *TV) On() {
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	fmt.Println("TV is OFF")
}

type Light struct {
}

func NewLight() *Light {
	return &Light{}
}

func (t *Light) On() {
	fmt.Println("Light is ON")
}

func (t *Light) Off() {
	fmt.Println("Light is OFF")
}

type Fan struct {
}

func (t *Fan) On() {
	fmt.Println("Fan is ON")
}

func (t *Fan) Off() {
	fmt.Println("Fan is OFF")
}
