package config

import (
	"encoding/json"
	"log"
	"os"
)

type AbstractConfig struct {
	Name		string
	Url 		string
	Interval	int
	ExpireSeconds	int
}

type BitckoinTickersConfig struct {
	BlockchainInfo	AbstractConfig
	CoindeskCom	AbstractConfig
}

type ExchangeTickersConfig struct {
	FixerIo		AbstractConfig
	Openexchange	AbstractConfig
}

type Configuration struct {
	BitcoinTickers	BitckoinTickersConfig
	ExchangeTickers	ExchangeTickersConfig
}

func GetConf() *Configuration {
	file, errF := os.Open("config.json")
	defer file.Close()
	if errF != nil {
		log.Panic(errF)
	}
	decoder := json.NewDecoder(file)
	conf := new(Configuration)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Panic(err)
	}
	return conf
}