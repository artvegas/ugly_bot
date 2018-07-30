package data

import (
	"time"

	gdax "github.com/preichenberger/go-gdax"
)

//client 	-> gdax.Client to connect to api
//ch		-> channel to write data to
func GetTicket(client gdax.Client, ch chan string) {
	for {
		stats, err := client.GetStats("BTC-USD")
		if err != nil {
			println(err.Error())
		}
		ch <- stats.Last
		time.Sleep(360 * time.Millisecond)
	}
}
