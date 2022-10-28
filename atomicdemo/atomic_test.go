package atomicdemo

import "testing"

func BenchmarkWithLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			IntStructWithLock()
		}
	}
}

func BenchmarkWithAtomic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			IntStructWithAtomic()
		}
	}
}
