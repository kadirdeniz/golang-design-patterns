package main

import (
	"design-patterns/creational-design-patterns/singleton-pattern/mutex"
	"design-patterns/creational-design-patterns/singleton-pattern/once"
)

func main() {
	once.GetOnceInstance()
	once.GetOnceInstance()
	mutex.GetMutexInstance()
	mutex.GetMutexInstance()
}
