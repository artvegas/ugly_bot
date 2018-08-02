package data

import (
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
)

func CollectTickerData() {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		println(err.Error())
	}

	subscribe := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			gdax.MessageChannel{
				Name: "heartbeat",
				ProductIds: []string{
					"BTC-USD",
				},
			},
			gdax.MessageChannel{
				Name: "level2",
				ProductIds: []string{
					"BTC-USD",
				},
			},
		},
	}
	if err := wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	for true {
		message := gdax.Message{}
		println(message.TradeId, message.TradeId)
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}

		if message.Type == "match" {
			println("Got a match")
			println(message.TradeId)
		}
	}
}
