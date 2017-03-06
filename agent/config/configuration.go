package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Configuration struct {
	stopRunning chan bool
}

var ConfFilePath string
var ConfFile *os.File
var running = make(chan bool)
var confFileModdTime time.Time
var confFileSize int64

type ConfObject struct {
	Object Collector
}

type Collector struct {
	server    string
	udpPort   string
	tcpPort   string
	instances []DBInstance
}

type DBInstance struct {
	server    string
	port      string
	user      string
	password  string
	slowquery string
}

/*
type (
	// for db instacne.
	dbInstance struct {
		server    string `json:"db.server"`
		port      string `json:"db.port"`
		user      string `json:"db.user"`
		password  string `json:"db.password"`
		slowquery string  `json:"db.slowquery.path"`
	}
	// for data collector.
	configure struct {
		configureData struct {
			//collectorIp      string       `json:"scouter.server"`
			//collectorUdpPort string       `json:"scouter.server.udp.port"`
			//collectorTcpPort string       `json:"scouter.server.tcp.port"`
			instances []dbInstance `json:"instacnes"`
		} `json:"configurations"`
	}
)
*/
func (conf *Configuration) load() {
	var confobj ConfObject
	fileInfo, e := ConfFile.Stat()
	if e != nil {
		//todo: error handling
	}

	if confFileModdTime != fileInfo.ModTime() || confFileSize != fileInfo.Size() {
		file, err := ioutil.ReadFile(ConfFilePath)
		if err != nil {
			//todo error
		}
		//json.NewDecoder(ConfFile).Decode(&confobj)
		json.Unmarshal(file, &confobj)
		confFileSize = fileInfo.Size()
		confFileModdTime = fileInfo.ModTime()
	}

}

func (conf *Configuration) init() bool {
	f, err := os.Open(ConfFilePath)
	if err != nil {
		return false
	} else {
		ConfFile = f
	}
	return true
}

func (conf *Configuration) Start() {
	go conf.run()
}

func (conf *Configuration) Stop() {
	conf.stopRunning <- true
}

func (conf *Configuration) run() {
	for {
		conf.load()
		time.Sleep(1 * time.Second)
		select {
		case <-conf.stopRunning:
			break
		default:
			continue
		}
	}
}

var instance *Configuration
var once sync.Once

// GetInstance returns configuraton singleton instance
func GetInstance() *Configuration {
	once.Do(func() {
		instance = &Configuration{}
		if instance.init() {
			instance.Start()
		}
	})
	return instance
}
