//file contains all actions we can perform to collect/save data
package data

import (
	"log"
	"time"

	gdax "github.com/preichenberger/go-gdax"
)

//client 	-> gdax.Client to connect to api
//ch		-> channel to write data to
func GetTickerData(client gdax.Client, ch chan *gdax.Ticker) {
	for {
		stats, err := client.GetTicker("BTC-USD")
		if err != nil {
			println(err.Error())
		}
		ch <- &stats
		time.Sleep(1000 * time.Millisecond)
	}
}

//get top 50 order book bids and ask
func GetOrderBook(client gdax.Client, ch chan *gdax.Book) {
	for {
		book, err := client.GetBook("BTC-USD", 2)
		if err != nil {
			log.Fatal(err)
		}
		ch <- &book
		time.Sleep(360 * time.Millisecond)
	}
}
