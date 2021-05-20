package designpattern

// 饿汉式
type Singleton struct{}

var singleton *Singleton

func InitSingleTon() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
