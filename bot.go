package main

import (
	"log"
	"net/http"
	"os"
	"time"

	actions "./actions"
	data "./data"
	ugly_bot_db "./database"
	_ "github.com/go-sql-driver/mysql"
	gdax "github.com/preichenberger/go-gdax"
)

func main() {
	// gets credentials for the api to work
	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	//waitFlg := false
	passphrase := os.Getenv("COINBASE_PASSPHRASE")

	//connect to the database first
	db, err := ugly_bot_db.ConnectToDB()
	//close connection after done using the program
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	// connect to coinbase api, returns client
	// which we can use to make calls to the
	// coinbase api
	client := gdax.NewClient(secret, key, passphrase)

	client.HttpClient = &http.Client{
		Timeout: 15 * time.Second,
	}

	ch := make(chan *gdax.Ticker)
	//orderBookCh := make(chan *gdax.Book)
	//boughtAt := 0.00
	//profitPercentage := 1.000003
	//lowRiskExitPercentage := 0.9999991
	go data.GetTickerData(*client, ch)
	//go data.GetOrderBook(*client, orderBookCh)
	//go data.CollectTickerData()
	// for i := range orderBookCh {
	// 	fmt.Println(i.Sequence, i.Asks, i.Bids)
	// }
	for i := range ch {
		actions.RecordPrice(i, db)
	}

}
