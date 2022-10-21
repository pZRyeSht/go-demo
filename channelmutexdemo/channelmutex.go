package channelmutex

import (
	"fmt"
	"time"
)

/*
使用 chan 实现互斥锁，至少有两种方式。
一种方式是先初始化一个 capacity 等于 1 的 Buffered Channel，然后再放入一个元素。这个元素就代表锁，谁取得了这个元素，就相当于获取了这把锁。
另一种方式是，先初始化一个 capacity 等于 1 的 Channel，它的“空槽”代表锁，谁能成功地把元素发送到这个 Channel，谁就获取了这把锁。
*/

// Mutex 使用channel实现互斥锁
type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{
		ch: make(chan struct{}, 1),
	}
	mu.ch <- struct{}{} // 方式一：初始化时放入一个元素
	return mu
}

// Lock 获取锁
func (m *Mutex) Lock() {
	<-m.ch
}

// Unlock 解锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock failed mutex")
	}
}

// TryLock 尝试获取锁
/*
tryLock() - 可轮询获取锁。如果成功，则返回 true；如果失败，则返回 false。
也就是说，这个方法无论成败都会立即返回，获取不到锁（锁已被其他线程获取）时不会一直等待。
*/
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// TryWithTimeout  超时获取锁
/*
tryLock(long, TimeUnit) - 可定时获取锁。和 tryLock() 类似，区别仅在于这个方法在获取不到锁时会等待一定的时间，
在时间期限之内如果还获取不到锁，就返回 false。如果如果一开始拿到锁或者在等待期间内拿到了锁，则返回 true。
*/
func (m *Mutex) TryWithTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

// IsLocked 锁是否已被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}


func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}
