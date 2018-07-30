package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	action "./actions"
	data "./data"
	ugly_bot_db "./database"
	_ "github.com/go-sql-driver/mysql"
	gdax "github.com/preichenberger/go-gdax"
)

func main() {
	// gets credentials for the api to work
	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	waitFlg := false
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

	ch := make(chan string)
	boughtAt := 0.00
	profitPercentage := 1.00003
	go data.GetTicket(*client, ch)
	for i := range ch {
		fmt.Println(i, time.Now())
		price, err := strconv.ParseFloat(i, 64)
		if err == nil && !waitFlg && price <= 8211.90 {
			boughtAt = price
			action.MakeOrder("1", db, i, "0.01")
			waitFlg = true
		}
		if err == nil && waitFlg && price >= boughtAt*profitPercentage {
			action.MakeOrder("2", db, i, "0.01")
			waitFlg = false
		}
	}

}
