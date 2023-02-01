package mutex

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

var mutexInstance *mutexSingleton

type mutexSingleton struct{}

func GetMutexInstance() *mutexSingleton {
	mutex.Lock()
	defer mutex.Unlock()
	if mutexInstance == nil {
		mutexInstance = new(mutexSingleton)
		fmt.Println("Creating unique instance now")
	}
	fmt.Println("Returning instance")
	return mutexInstance
}
