package chapter5_singleton

import (
	"fmt"
	"math/rand"
	"sync"
)

type OnlyOneEver struct {
	id int
}

var instance *OnlyOneEver

var once = &sync.Once{}

func GetOnlyOneEver() *OnlyOneEver {
	once.Do(func() {
		if instance == nil {
			id := rand.Intn(100)
			fmt.Println("Created instance", id)
			instance = &OnlyOneEver{id}
		}
	})

	return instance
}
