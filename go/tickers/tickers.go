package tickers

import (
	"encoding/json"
	"net/http"
	"time"
)

type AbstractBitcoinExchange struct {
	Usd float32
	Eur float32
}

type BlockchainJson struct {
	Usd struct {
		Sell float32
	}
	Eur struct {
		Sell float32
	}
}

type CoindeskJson struct {
	Bpi struct {
		Usd struct {
			RateFloat float32
		}
		Eur struct {
			RateFloat float32
		}
	}
}

type FixerIoJson struct {
	Rates struct {
		Usd float32
	}
}
type OpenexchangeJson struct {
	Rates struct {
		Eur float32
	}
}

type AbstractTicker struct {
	Name            string
	Url             string
	Active          bool
	ExpireTime      time.Duration
	RequestInterval time.Duration
	requestTicker   *time.Ticker
	ResultChan      chan bool
	LastSuccessReq  time.Time
	JsonStruct      interface{}
	Parse           func(interface{}) map[string]float32
}

type AbstractBitcoinTicker struct {
	AbstractTicker
}

type AbstractExchangeTicker struct {
	AbstractTicker
}

type AbstractTickerInterface interface {
	makeRequest()
	schedule()
	Start()
	Parse()
	Stop()
}

func (at *AbstractTicker) Start() {
	at.schedule()
	go at.makeRequest()
}

func (at *AbstractTicker) Stop() {
	at.requestTicker.Stop()
}

func (at *AbstractTicker) makeRequest() {
	response, err := http.Get(at.Url)
	defer response.Body.Close()

	if err != nil {
		at.Active = false
		//log.Fatal(err)
		at.ResultChan <- false
		return
	}
	decoder := json.NewDecoder(response.Body)

	errD := decoder.Decode(at.JsonStruct)
	if errD != nil {
		at.Active = false
		at.ResultChan <- false
		return
	}
	at.LastSuccessReq = time.Now()
	at.Active = true
	at.ResultChan <- true
}

func (at *AbstractTicker) schedule() {
	at.requestTicker = time.NewTicker(at.RequestInterval)
	go func() {
		for t := range at.requestTicker.C {
			_ = t
			at.makeRequest()
		}
	}()
}

func BlockchainTickerParse(data interface{}) map[string]float32 {

	return map[string]float32{
		"USD": data.(*BlockchainJson).Usd.Sell,
		"EUR": data.(*BlockchainJson).Eur.Sell,
	}
}

func CoindeskTickerParse(data interface{}) map[string]float32 {
	return map[string]float32{
		"USD": data.(*CoindeskJson).Bpi.Usd.RateFloat,
		"EUR": data.(*CoindeskJson).Bpi.Eur.RateFloat,
	}
}

func FixerIoExchangeTickerParse(data interface{}) map[string]float32 {
	return map[string]float32{
		"EUR/USD": data.(*FixerIoJson).Rates.Usd,
	}
}

func OpenexchangeTickerParse(data interface{}) map[string]float32 {
	return map[string]float32{
		"EUR/USD": 1 / data.(*OpenexchangeJson).Rates.Eur,
	}
}
