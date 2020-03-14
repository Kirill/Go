package test1

import "testing"

func BenchmarkF(b *testing.B) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	bases := []int{0, 224, 256, 384, 512}
	for i := 0; i < b.N; i++ {
		f(s, bases[i%len(bases)])
	}
}
