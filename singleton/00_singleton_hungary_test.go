package designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton_GetInstance(t *testing.T) {
	InitSingleTon()

	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	InitSingleTon()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Fatalf("not equal")
			}
		}
	})
}
