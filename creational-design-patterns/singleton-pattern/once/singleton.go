package once

import (
	"fmt"
	"sync"
)

var once sync.Once

var onceInstance *onceSingleton

type onceSingleton struct{}

func GetOnceInstance() *onceSingleton {
	once.Do(func() {
		fmt.Println("Creating unique instance now")
		onceInstance = new(onceSingleton)
	})
	fmt.Println("Returning instance")
	return onceInstance
}
