package channelmutex

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}
