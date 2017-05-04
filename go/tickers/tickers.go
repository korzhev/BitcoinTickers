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

type ATInterface interface {
	Parse()
}

type AbstractTicker struct {
	Name            string
	Url             string
	Active          bool
	ExpireTime      time.Duration
	RequestInterval time.Duration
	StopChan        chan bool
	ResultChan      *chan bool
	LastSuccessReq  time.Time
	JsonStruct	interface{}
}

type AbstractBitcoinTicker struct {
	AbstractTicker
	LastData        AbstractBitcoinExchange
}

type AbstractExchangeTicker struct {
	AbstractTicker
	LastData        float32
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
	makeRequest() ()
	schedule() ()
	Parse()
	Start()
	Stop()
}
func Parse (ati AbstractTickerInterface) {
	return ati.Parse()
}

func (at AbstractTicker) Parse()  {
	log.Print(at)
}

func (at AbstractTicker) Start() {
	at.schedule()
}

func (at AbstractTicker) Stop()  {
	at.StopChan <- true
}

func (at AbstractTicker) makeRequest() {
	response, err := http.Get(at.Url)
	if err != nil {
		at.Active = false
		//log.Fatal(err)
		*at.ResultChan <- false
		return
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)

	errD := decoder.Decode(at.JsonStruct)
	if errD != nil {
		at.Active = false
		//log.Fatal(errD)
		*at.ResultChan <- false
		return
	}
	at.LastSuccessReq = time.Now()
	at.Active = true
}

func (at AbstractTicker) getData() {
	at.makeRequest()
	at.Parse()
}

func (at AbstractTicker) schedule() {
	at.StopChan = make(chan bool)
	go func() {
		for {
			at.getData()
			select {
			case <-time.After(at.RequestInterval):
			case <-at.StopChan:
				return
			}
		}
	}()
}

func (bt BlockchainTicker) Parse() {
	bt.LastData = AbstractBitcoinExchange{
		Usd: 	bt.JsonStruct.Usd.Sell,
		Eur: 	bt.JsonStruct.Eur.Sell,
	}
}

func (ct CoindeskTicker) Parse() {
	ct.LastData = AbstractBitcoinExchange{
		Usd: 	ct.JsonStruct.Bpi.Usd.RateFloat,
		Eur: 	ct.JsonStruct.Bpi.Eur.RateFloat,
	}
}

func (ft FixerIoExchangeTicker) Parse() {
	ft.LastData = ft.JsonStruct.Rates.Usd
}

func (ot OpenexchangeTicker) Parse() {
	ot.LastData = 1 / ot.JsonStruct.Rates.Eur
}


