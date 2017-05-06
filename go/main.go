package main

import (
	"./config"
	"./tickers"
	"fmt"
	"time"
)

var conf = config.GetConf()

type printData struct {
	USD                   float32
	EUR                   float32
	EURUSD                float32
	ActiveBitcoinTickers  int
	ActiveCurrencyTickers int
}

func processActiveBitcoinTickers(list []*tickers.AbstractBitcoinTicker) (map[string]float32, int) {
	counter := 0
	min := float32(0.0)
	result := map[string]float32{"USD": 0.0, "EUR": 0.0}
	now := time.Now()
	for _, t := range list {
		if t.Active {
			counter++
		}
		if t.LastSuccessReq.Add(t.ExpireTime).After(now) {
			data := t.Parse(t.JsonStruct)
			if min == 0.0 || min < data["EUR"] {
				min = data["EUR"]
				result = data
			}
		}
	}
	return result, counter
}

func processActiveExchangeTickers(list []*tickers.AbstractExchangeTicker) (map[string]float32, int) {
	counter := 0
	min := float32(0.0)
	result := map[string]float32{"EUR/USD": 0.0}
	now := time.Now()
	for _, t := range list {
		if t.Active {
			counter++
		}
		if t.LastSuccessReq.Add(t.ExpireTime).After(now) {
			data := t.Parse(t.JsonStruct)
			if min == 0.0 || min < data["EUR/USD"] {
				min = data["EUR/USD"]
				result = data
			}
		}
	}
	return result, counter
}

func main() {
	//messages := make(chan bool)
	BitcoinTicker1 := tickers.AbstractBitcoinTicker{
		AbstractTicker: tickers.AbstractTicker{
			Name:            conf.BitcoinTickers.BlockchainInfo.Name,
			Url:             conf.BitcoinTickers.BlockchainInfo.Url,
			ExpireTime:      time.Second * time.Duration(conf.BitcoinTickers.BlockchainInfo.ExpireTime),
			RequestInterval: time.Second * time.Duration(conf.BitcoinTickers.BlockchainInfo.Interval),
			ResultChan:      make(chan bool),
			LastSuccessReq:  time.Now(),
			JsonStruct:      new(tickers.BlockchainJson),
			Parse:           tickers.BlockchainTickerParse,
		},
	}
	BitcoinTicker2 := tickers.AbstractBitcoinTicker{
		AbstractTicker: tickers.AbstractTicker{
			Name:            conf.BitcoinTickers.CoindeskCom.Name,
			Url:             conf.BitcoinTickers.CoindeskCom.Url,
			ExpireTime:      time.Second * time.Duration(conf.BitcoinTickers.CoindeskCom.ExpireTime),
			RequestInterval: time.Second * time.Duration(conf.BitcoinTickers.CoindeskCom.Interval),
			ResultChan:      make(chan bool),
			LastSuccessReq:  time.Now(),
			JsonStruct:      new(tickers.CoindeskJson),
			Parse:           tickers.CoindeskTickerParse,
		},
	}

	//ExchangeTickerList := []*tickers.AbstractExchangeTicker{
	//	{
	//		AbstractTicker: tickers.AbstractTicker{
	//			Name:            conf.ExchangeTickers.FixerIo.Name,
	//			Url:             conf.ExchangeTickers.FixerIo.Url,
	//			ExpireTime:      time.Second * time.Duration(conf.ExchangeTickers.FixerIo.ExpireTime),
	//			RequestInterval: time.Second * time.Duration(conf.ExchangeTickers.FixerIo.Interval),
	//			ResultChan:      make(chan bool),
	//			LastSuccessReq:  time.Now(),
	//			JsonStruct:      new(tickers.FixerIoJson),
	//			Parse:           tickers.FixerIoExchangeTickerParse,
	//		},
	//	},
	//	{
	//		AbstractTicker: tickers.AbstractTicker{
	//			Name:            conf.ExchangeTickers.Openexchange.Name,
	//			Url:             conf.ExchangeTickers.Openexchange.Url,
	//			ExpireTime:      time.Second * time.Duration(conf.ExchangeTickers.Openexchange.ExpireTime),
	//			RequestInterval: time.Second * time.Duration(conf.ExchangeTickers.Openexchange.Interval),
	//			ResultChan:      make(chan bool),
	//			LastSuccessReq:  time.Now(),
	//			JsonStruct:      new(tickers.OpenexchangeJson),
	//			Parse:           tickers.OpenexchangeTickerParse,
	//		},
	//	},
	//}
	//for _, t := range BitcoinTickerList {
	//	t.Start()
	//}
	//for _, t := range ExchangeTickerList {
	//	t.Start()
	//}

	BitcoinTicker1.Start()
	BitcoinTicker2.Start()
	//str := "BTC/USD: %.2f EUR/USD: %.2f BTC/EUR: %.2f Active sources: BTC/USD %d of %d EUR/USD %d of %d Time: %v\n"
	for {
		select {
		case msg1 := <-BitcoinTicker1.ResultChan:
			fmt.Println("received", msg1)
		case msg2 := <-BitcoinTicker2.ResultChan:
			fmt.Println("received", msg2)
		}
	}

	//for msg := range messages {
	//	_ = msg
	//	bestBitcoin, activeBitcoinTickers := processActiveBitcoinTickers(BitcoinTickerList)
	//	bestCurrency, activeExchangeTickers := processActiveExchangeTickers(ExchangeTickerList)
	//	fmt.Printf(str,
	//		bestBitcoin["USD"],
	//		bestCurrency["EUR/USD"],
	//		bestBitcoin["EUR"],
	//		activeBitcoinTickers,
	//		len(BitcoinTickerList),
	//		activeExchangeTickers,
	//		len(ExchangeTickerList),
	//		time.Now())
	//}
}
