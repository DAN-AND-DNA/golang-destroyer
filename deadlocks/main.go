package main

import (
	//"sync"
	sync "github.com/sasha-s/go-deadlock"
	"time"
)

func main() {
	recursiveLock1()
	recursiveLock2()
}

// recursive locking (递归加锁)
func recursiveLock1() {
	var m sync.Mutex

	func() {
		m.Lock()
		defer m.Unlock()

		m.Lock()
	}()

}

// recursive locking (递归加锁)
func recursiveLock2() {
	var m sync.RWMutex

	go func() {
		for {
			// FIXME 默认的情况：只有全部协程都死锁了才会panic，这里是为了模拟死锁协程泄露的情况
			// "github.com/sasha-s/go-deadlock" 发现死锁会直接panic
			time.Sleep(100 * time.Second)
		}
	}()

	go func() {
		m.RLock()
		defer m.RUnlock()
		{
			time.Sleep(3 * time.Second)
			m.RLock() // 存在“阻塞状态的writer”，则新的reader会被阻塞
			defer m.RUnlock()
		}
	}()

	time.Sleep(1 * time.Second)
	m.Lock() // 阻塞状态的writer
	defer m.Unlock()
}
