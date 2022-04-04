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

var lock = &sync.Mutex{}

func GetOnlyOneEver() *OnlyOneEver {

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			id := rand.Intn(100)
			fmt.Println("Created instance", id)
			instance = &OnlyOneEver{id}
		}
	}
	return instance
}
