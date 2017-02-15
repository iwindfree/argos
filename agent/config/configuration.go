package config

import (
	"sync"
)

type Configuraton struct {
}

type (
	// for db instacne
	dbInstance struct {
		server   string `json:"db.server"`
		port     string `json:"db.port"`
		user     string `json:"db.user"`
		password string `json:"db.password"`
	}
	// for configuration root
	configure struct {
		configureData struct {
			collectorIp      string       `json:"scouter.server"`
			collectorUdpPort string       `json:"scouter.server.udp.port"`
			collectorTcpPort string       `json:"scouter.server.tcp.port"`
			instances        []dbInstance `json:"instacnes"`
		} `json:"configurations"`
	}
)

var instance *Configuraton
var once sync.Once

func GetInstance() *Configuraton {
	once.Do(func() {
		instance = &Configuraton{}
	})
	return instance
}
