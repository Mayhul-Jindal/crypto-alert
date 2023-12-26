package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"nhooyr.io/websocket"
)

type SubscribeResponse struct {
	Result interface{} `json:"result"`
	Id     int         `json:"id"`
}

type StreamData struct {
	Price string `json:"p"`
}

type StreamResponse struct {
	Stream string     `json:"stream"`
	Data   StreamData `json:"data"`
}

type cryptoWatcher struct {
	market map[currency]float64
	ws     *websocket.Conn
	errch  chan error

	// throttling my market readers for demo purposes
	ticker *time.Ticker
}

func NewCryptoWatcher(ctx context.Context, currencies []currency) (*cryptoWatcher, error) {
	c, _, err := websocket.Dial(ctx, "wss://stream.binance.com/stream", nil)
	if err != nil {
		return nil, err
	}

	// map init
	market := make(map[currency]float64)
	for _, curr := range currencies {
		market[curr] = 0
	}

	// prepring subscribe request payload
	subscribePayload := map[string]interface{}{
		"method": "SUBSCRIBE",
		"params": currencies,
		"id":     1,
	}
	payloadBytes, err := json.Marshal(subscribePayload)
	if err != nil {
		return nil, err
	}

	// sending subscribe request
	err = c.Write(ctx, websocket.MessageText, payloadBytes)
	if err != nil {
		return nil, err
	}

	// reading subscribe response and checking if subscription was successful
	_, p, err := c.Read(ctx)
	if err != nil {
		return nil, err
	}
	var pubResponse SubscribeResponse
	err = json.Unmarshal(p, &pubResponse)
	if err != nil {
		return nil, err
	}
	log.Println(pubResponse)
	if pubResponse.Result != nil {
		return nil, ErrSubscriptionFailed
	}

	return &cryptoWatcher{
		market: market,
		ws:     c,
		errch:  make(chan error),
		ticker: time.NewTicker(100 * time.Millisecond),
	}, nil
}

func (c *cryptoWatcher) Run(ctx context.Context) error {
	// todo unmarshall and fill the market
	go c.fillMarket(ctx)

	// start comparing with target price of users
	for curr := range c.market {
		go c.startComparing(curr)
	}

	// handles errors, can be a potential centalized thingy
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-c.errch:
			log.Println(err)
		}
	}
}

func (c *cryptoWatcher) fillMarket(ctx context.Context) {
	for {
		_, p, err := c.ws.Read(ctx)
		if err != nil {
			c.errch <- err
		}

		// unmarshall and fill the market
		var streamResponse StreamResponse
		err = json.Unmarshal(p, &streamResponse)
		if err != nil {
			c.errch <- err
		}

		price, err := strconv.ParseFloat(streamResponse.Data.Price, 64)
		if err != nil {
			c.errch <- err
		}

		c.market[currency(streamResponse.Stream)] = price
	}
}

func (c *cryptoWatcher) startComparing(curr currency) {
	// reaading market price after tick time
	for range c.ticker.C {
		switch c.market[curr] {
		case 0:
			continue

		default:
			log.Println(curr, c.market[curr])
			// todo compare with target price of users
		}
	}
}