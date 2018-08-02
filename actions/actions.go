// file contains all actions that a bot can do

package actions

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/preichenberger/go-gdax"
)

//puts virtual order into database
func MakeOrder(order_type string, db *sql.DB, price string, quantity string) {
	order_type_print := ""
	stmt, err := db.Prepare("INSERT INTO order_book(order_type_id, offer_asset_type_id, want_asset_type_id, quantity, price, order_status_id, original_order_id, market_order, expiration_date, date_modified, date_created) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	switch order_type {
	case "1":
		order_type_print = "BUY"
		//order_type_id = 1 -> buy_order
		//offer_asset_type_id = 1 -> USD
		//want_asset_type_id = 2 -> BTC
		//order_status_id = 1 -> pending
		_, err := stmt.Exec(order_type, "1", "2", quantity, price, "1", "NULL", "NULL", time.Now().AddDate(0, 0, 0), time.Now().Local(), time.Now().Local())
		if err != nil {
			log.Fatal(err)
		}
		break
	case "2":
		order_type_print = "SELL"
		//order_type_id = 1 -> buy_order
		//offer_asset_type_id = 1 -> USD
		//want_asset_type_id = 2 -> BTC
		//order_status_id = 1 -> pending
		_, err := stmt.Exec(order_type, "2", "1", quantity, price, "1", "NULL", "NULL", time.Now().AddDate(0, 0, 0), time.Now().Local(), time.Now().Local())
		if err != nil {
			log.Fatal(err)
		}

		break
	}

	result := fmt.Sprintf("Added order: TYPE [%s ORDER] QUANITY [%s BTC] PRICE [%s USD]", order_type_print, quantity, price)
	fmt.Println(result)
}

func RecordPrice(ticker *gdax.Ticker, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO btc_usd_history(
		trade_id, price, size,
		 time, bid, ask, volume) 
		 VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	string_time := ticker.Time.Time().String()
	result, err := stmt.Exec(
		ticker.TradeId,
		ticker.Price,
		ticker.Size,
		string_time,
		ticker.Bid,
		ticker.Ask,
		ticker.Volume,
	)
	if err != nil {
		log.Fatal(err)
	}
	println(result.LastInsertId)
}
