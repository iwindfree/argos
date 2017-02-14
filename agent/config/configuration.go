package config

import (
	"sync"
)

type Configuraton struct {
}

var instance *Configuraton
var once sync.Once

func GetInstance() *Configuraton {
	once.Do(func() {
		instance = &Configuraton{}
	})
	return instance
}
