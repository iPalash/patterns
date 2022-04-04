package chapter5_singleton

import (
	"math/rand"
)

type OnlyOneEver struct {
	id int
}

var instance *OnlyOneEver

func GetOnlyOneEver() *OnlyOneEver {
	if instance == nil {
		instance = &OnlyOneEver{rand.Intn(100)}
	}
	return instance
}
