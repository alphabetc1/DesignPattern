package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data int
}

var sin *singleton

var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		sin = &singleton{15}
	})

	fmt.Println("singleton: ", sin, &sin)
	return sin
}
