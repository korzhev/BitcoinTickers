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

type AbctractTicker struct {
	Name            string
	url             string
	Active          bool
	expireTime      time.Duration
	requestInterval time.Duration
	LastData        AbstractBitcoinExchange
	stopChan        chan bool
	LastSuccessReq  time.Time
}

type BlockchainTicker struct {
	AbctractTicker
}

func (at AbctractTicker) Parse()  {
	log.Fatalf(".Parse() not implemented at " + at.Name)
}

func (at AbctractTicker) Stop()  {
	at.stopChan <- true
}

func (at AbctractTicker) makeRequest(url string) (*json.Decoder, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		decoder := json.NewDecoder(response.Body)
		return decoder, nil
	}
}

func (at AbctractTicker) schedule(function func(), delay time.Duration) (chan bool) {
	at.stopChan = make(chan bool)
	go func() {
		for {
			function()
			select {
			case <-time.After(delay):
			case <-at.stopChan:
				return
			}
		}
	}()
	return at.stopChan
}

func (bt BlockchainTicker) Parse(decoder *json.Decoder) {
	jsonData := new(BlockchainJson)
	err := decoder.Decode(&jsonData)
	if err != nil {
		log.Fatal(err)
	}
	bt.LastSuccessReq = time.Now()
	bt.LastData = AbstractBitcoinExchange{Usd: jsonData.Usd.Sell, Eur:jsonData.Eur.Sell }
}

type AbstractTickerInterface interface {
	MakeRequest(url string) (*json.Decoder, error)
	Parse()
	//Parse() (map[string]float32)
}



