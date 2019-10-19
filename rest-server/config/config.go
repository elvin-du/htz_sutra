package config

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	configPath = "./config.json"

	logFile = "./log.txt"
)

var (
	DefaultConfig conf
)

type Log struct {
	Output string `json:"output"`
	Level  string `json:"level"`
}

type conf struct {
	DBAddress   string `json:"db_address"`
	DBUser      string `json:"db_user"`
	DBPassword  string `json:"db_password"`
	DBName      string `json:"db_name"`
	NodeAddress string `json:"node_address"`
	LcdAddress  string `json:"lcd_address"`
	HTTPAddress string `json:"http_address"`
	MongoURI    string
	Log         Log `json:"log"`
}

func init() {
	f, err := os.Open(configPath)
	if nil != err {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&DefaultConfig)
	if nil != err {
		panic(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", DefaultConfig.DBUser, DefaultConfig.DBPassword, DefaultConfig.DBAddress)
	log.Println("mongoURI:", mongoURI)
	DefaultConfig.MongoURI = mongoURI

	if "file" == DefaultConfig.Log.Output {
		f, err := os.OpenFile(logFile,os.O_RDWR|os.O_CREATE, 0666)
		if nil != err {
			panic(err)
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}

	level, err := log.ParseLevel(DefaultConfig.Log.Level)
	if nil != err {
		panic(err)
	}

	log.SetLevel(level)
}
