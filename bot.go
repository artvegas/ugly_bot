package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	gdax "github.com/preichenberger/go-gdax"
)

func getTicket(client gdax.Client, ch chan string) {
	for {
		stats, err := client.GetStats("BTC-USD")
		if err != nil {
			println(err.Error())
		}
		ch <- stats.Last
		time.Sleep(360 * time.Millisecond)
	}
}

func makeBuyOrder(price string) {
	fmt.Println("Making buy order at: ", price)
}

func makeSellOrder(price string) {
	fmt.Println("Making sell order at: ", price)
}

func main() {
	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	waitFlg := false
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	client := gdax.NewClient(secret, key, passphrase)

	client.HttpClient = &http.Client{
		Timeout: 15 * time.Second,
	}

	ch := make(chan string)
	go getTicket(*client, ch)
	for i := range ch {
		fmt.Println(i, time.Now())
		price, err := strconv.ParseFloat(i, 64)
		if err == nil && !waitFlg && price <= 8249.99 {
			makeBuyOrder(i)
			waitFlg = true
		}
		if err == nil && waitFlg && price >= 8250.00 {
			makeSellOrder(i)
			waitFlg = false
		}
	}

}
