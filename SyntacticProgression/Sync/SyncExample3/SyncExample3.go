package main

import "sync"

type Config struct{}

var instance *Config
var once sync.Once

/*
只有在第一次调用InitConfig()获取Config 指针的时候才会执行once.Do(func(){instance = &Config{
})语句，执行完之后instance就驻留在内存中，后面再次执行InitConfig()的时候，就直接返回内存中
的instance。
*/

func InitConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func main() {

}
