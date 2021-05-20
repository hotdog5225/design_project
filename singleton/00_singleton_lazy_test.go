package designpattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingletonLazyInstance(t *testing.T) {
	assert.Equal(t, GetSingletonLazyInstance(), GetSingletonLazyInstance())
}

func BenchmarkGetSingletonLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingletonLazyInstance() != GetSingletonLazyInstance() {
				b.Fatalf("not equal!")
			}
		}
	})
}
