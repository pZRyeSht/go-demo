package atomicdemo

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
atomic包与读写锁的并发安全基准测试
**/

type intStruct struct {
	q []int
}

// IntStructWithLock SetIntStructWithLock 锁读写
func IntStructWithLock() {
	lock := sync.RWMutex{}
	is := intStruct{}
	// 写
	go func() {
		i := 0
		for {
			i++
			lock.Lock()
			is.q = []int{i, i+1, i+3}
			lock.Unlock()
		}
	}()
	
	// 读
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				lock.RLock()
				fmt.Printf("#%v\n", is)
				lock.RUnlock()
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

// IntStructWithAtomic atomic.Value
func IntStructWithAtomic() {
	ato := atomic.Value{}
	is := intStruct{}
	// 写
	go func() {
		i := 0
		for {
			i++
			is.q = []int{i, i+1, i+3}
			ato.Store(is)
		}
	}()

	// 读
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				temp := ato.Load()
				fmt.Printf("#%v\n", temp)
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
