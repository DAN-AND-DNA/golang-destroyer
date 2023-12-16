package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	doubleUnlock()
	unlockByOther()
}

func doubleUnlock() {
	var m sync.Mutex
	m.Lock()
	m.Unlock()
	m.Unlock()
}

func unlockByOther() {
	var m sync.Mutex

	m.Lock()

	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("first")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		m.Unlock()
	}()

	time.Sleep(3 * time.Second)
}
