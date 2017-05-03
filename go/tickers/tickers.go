package tickers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AbstractBitcoinExchange struct {
	Usd	float32
	Eur	float32
}

type BlockchainJson struct {
	Usd struct {
		Sell	float32
	}
	Eur struct {
		Sell	float32
	}
}

type CoindeskJson struct {
	Bpi	struct {
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
		Usd	float32
	}
}
type OpenexchangeJson struct {
	Rates struct {
		Eur	float32
	}
}

type AbstractTicker struct {
	Name            string
	url             string
	Active          bool
	expireTime      time.Duration
	requestInterval time.Duration
	stopChan        chan bool
	resultChan	chan bool
	LastSuccessReq  time.Time
}

type AbstractBitcoinTicker struct {
	AbstractTicker
	LastData        AbstractBitcoinExchange
	JsonStruct	interface{}
}

type AbstractExchangeTicker struct {
	AbstractTicker
	LastData        float32
	LastSuccessReq  time.Time
}

type BlockchainTicker struct {
	AbstractBitcoinTicker
	JsonStruct	BlockchainJson
}

type CoindeskTicker struct {
	AbstractBitcoinTicker
	JsonStruct	CoindeskJson
}

type FixerIoExchangeTicker struct {
	AbstractExchangeTicker
	JsonStruct	FixerIoJson
}

type OpenexchangeTicker struct {
	AbstractExchangeTicker
	JsonStruct	OpenexchangeJson
}

type AbstractTickerInterface interface {
	makeRequest(url string) (*json.Decoder, error)
	schedule(func(), time.Duration) (chan bool)
	Parse()
	Start()
	Stop()
}

func (at AbstractBitcoinTicker) Parse()  {
	log.Fatalf(".Parse() not implemented at " + at.Name)
}

func (at AbstractBitcoinTicker) Start() {
	at.schedule()
}

func (at AbstractBitcoinTicker) Stop()  {
	at.stopChan <- true
}

func (at AbstractBitcoinTicker) makeRequest() {
	response, err := http.Get(at.url)
	if err != nil {
		at.Active = false
		//log.Fatal(err)
		at.resultChan <- false
		return
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)

	errD := decoder.Decode(at.JsonStruct)
	if errD != nil {
		at.Active = false
		//log.Fatal(errD)
		at.resultChan <- false
		return
	}
	at.LastSuccessReq = time.Now()
	at.Active = true
}

func (at AbstractBitcoinTicker) getData() {
	at.makeRequest()
	at.Parse()
}

func (at AbstractBitcoinTicker) schedule() {
	at.stopChan = make(chan bool)
	go func() {
		for {
			at.getData()
			select {
			case <-time.After(at.requestInterval):
			case <-at.stopChan:
				return
			}
		}
	}()
}

func (bt BlockchainTicker) Parse() {
	bt.LastData = AbstractBitcoinExchange{
		Usd: bt.JsonStruct.Usd.Sell,
		Eur:bt.JsonStruct.Eur.Sell,
	}
}

func (ct CoindeskTicker) Parse() {
	ct.LastData = AbstractBitcoinExchange{
		Usd: ct.JsonStruct.Bpi.Usd.RateFloat,
		Eur: ct.JsonStruct.Bpi.Eur.RateFloat,
	}
}

func (ft FixerIoExchangeTicker) Parse() {
	ft.LastData = ft.JsonStruct.Rates.Usd
}

func (ot OpenexchangeTicker) Parse() {
	ot.LastData = 1 / ot.JsonStruct.Rates.Eur
}


