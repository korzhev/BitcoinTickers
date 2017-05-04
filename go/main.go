package main

import (
	"./config"
	"./tickers"
	"time"
	"log"
)
var conf = config.GetConf()

func main() {
	message:= make(chan bool)

	BT:= &tickers.BlockchainTicker{
		AbstractBitcoinTicker:tickers.AbstractBitcoinTicker{
			AbstractTicker: tickers.AbstractTicker{
				Name: conf.BitcoinTickers.BlockchainInfo.Name,
				Url: conf.BitcoinTickers.BlockchainInfo.Url,
				Active:false,
				ExpireTime: 60 * time.Second,
				RequestInterval: 600 * time.Second,
				StopChan: make(chan bool),
				ResultChan: &message,
				LastSuccessReq: time.Now(),
			},
			LastData: *new(tickers.AbstractBitcoinExchange),
		},
		JsonStruct: *new(tickers.BlockchainJson),
	}

	CT:= &tickers.CoindeskTicker{
		AbstractBitcoinTicker: tickers.AbstractBitcoinTicker{
			AbstractTicker: tickers.AbstractTicker{
				Name: conf.BitcoinTickers.CoindeskCom.Name,
				Url: conf.BitcoinTickers.CoindeskCom.Url,
				Active:false,
				ExpireTime: 60 * time.Second,
				RequestInterval: 600 * time.Second,
				StopChan: make(chan bool),
				ResultChan: &message,
				LastSuccessReq: time.Now(),
			},
			LastData: *new(tickers.AbstractBitcoinExchange),
		},
		JsonStruct: *new(tickers.CoindeskJson),
	}

	OT:= &tickers.OpenexchangeTicker{
		AbstractExchangeTicker: tickers.AbstractExchangeTicker{
			AbstractTicker: tickers.AbstractTicker{
				Name: conf.ExchangeTickers.Openexchange.Name,
				Url: conf.ExchangeTickers.Openexchange.Url,
				Active:false,
				ExpireTime: 60 * time.Second,
				RequestInterval: 600 * time.Second,
				StopChan: make(chan bool),
				ResultChan: &message,
				LastSuccessReq: time.Now(),
			},
			LastData: 0.0,
		},
		JsonStruct: *new(tickers.OpenexchangeJson),
	}

	FT:= &tickers.FixerIoExchangeTicker{
		AbstractExchangeTicker: tickers.AbstractExchangeTicker{
			AbstractTicker: tickers.AbstractTicker{
				Name: conf.ExchangeTickers.FixerIo.Name,
				Url: conf.ExchangeTickers.FixerIo.Url,
				Active:false,
				ExpireTime: 60 * time.Second,
				RequestInterval: 600 * time.Second,
				StopChan: make(chan bool),
				ResultChan: &message,
				LastSuccessReq: time.Now(),
			},
			LastData: 0.0,
		},
		JsonStruct: *new(tickers.FixerIoJson),
	}
	BT.Start()
	CT.Start()
	FT.Start()
	OT.Start()
	log.Print(BT.LastData)
	log.Print(CT.LastData)
	log.Print(FT.LastData)
	log.Print(OT.LastData)

	for msg:= range message {
		log.Print(msg)
	}

	//for range message {
	//	log.Print(BT.LastData)
	//	log.Print(CT.LastData)
	//	log.Print(FT.LastData)
	//	log.Print(OT.LastData)
	//}

}
