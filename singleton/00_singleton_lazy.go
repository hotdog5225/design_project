package designpattern

import "sync"

type SingletonLazy struct {
}

var (
	singletonLazy *SingletonLazy
	once = sync.Once{}
)

func GetSingletonLazyInstance() *SingletonLazy {
	if singletonLazy == nil {
		once.Do(func() {
			singletonLazy = &SingletonLazy{}
		})
	}
	return singletonLazy
}