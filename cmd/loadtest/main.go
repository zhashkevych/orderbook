package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gosuri/uilive"
	"github.com/zhashkevych/orderbook/orderbook"
)

// TODO: create order for pair (USDT-BTC for example)
// TODO: add update / remove order (USDT-BTC for example)
// TODO: set requests amount (1 mil for example) & collect stats to file
// TODO: gRPC

const (
	URL_BASE = "http://localhost:8080/"

	MIN_PRICE = 1
	MAX_PRICE = 50

	MIN_AMOUNT = 1
	MAX_AMOUNT = 50000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	createOrderLatency := make(chan time.Duration, 1000)
	createdOrders, createOrderErrors := 0, 0

	// create order requests with random interval
	go func(requests, errors *int) {
		for {
			go func(errors *int) {
				if err := createOrder(orderbook.NewOrder(
					rand.Intn(MAX_PRICE-MIN_PRICE)+MIN_PRICE,
					rand.Intn(MAX_AMOUNT-MIN_AMOUNT)+MIN_AMOUNT,
					orderbook.OrderType(rand.Intn(2)),
				), createOrderLatency); err != nil {
					*errors++
				}
			}(errors)

			*requests++

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}
	}(&createdOrders, &createOrderErrors)

	// listOrderbookLatency := make(chan time.Duration, 1000)
	// listOrderbookErrros, listOrderbookRequests := 0, 0
	// // get orderbook requests with random interval
	// go func(requests, errors *int) {
	// 	for {
	// 		go func(errors *int) {
	// 			_, err := listOrderbook(listOrderbookLatency)
	// 			if err != nil {
	// 				*errors++
	// 			}

	// 			*requests++

	// 			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	// 		}(errors)

	// 	}
	// }(&listOrderbookRequests, &listOrderbookErrros)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	writer := uilive.New()
	writer.Start()

	go func() {
		totalLatency := time.Duration(0)
		for lat := range createOrderLatency {
			totalLatency += lat

			avg := totalLatency / time.Duration(createdOrders)

			fmt.Fprintf(writer, "[CREATE ORDER] average latency / created orders / errors (%d ns / %d / %d) \n",
				avg.Nanoseconds(), createdOrders, createOrderErrors)
		}
	}()

	// go func() {
	// 	totalLatency := time.Duration(0)
	// 	for lat := range listOrderbookLatency {
	// 		totalLatency += lat

	// 		avg := totalLatency / time.Duration(listOrderbookRequests)

	// 		fmt.Fprintf(writer, "[LIST ORDERBOOK] average latency / created orders / errors (%d ns / %d / %d) \n",
	// 			avg.Nanoseconds(), createdOrders, createOrderErrors)
	// 	}
	// }()

	<-sigint

	fmt.Fprintln(writer, "Stopped")
	writer.Stop() // flush and stop rendering
}

func createOrder(inp *orderbook.Order, latency chan time.Duration) error {
	body, err := json.Marshal(inp)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%sorders/ask", URL_BASE)
	if inp.OrderType == orderbook.TYPE_BID {
		url = fmt.Sprintf("%sorders/bid", URL_BASE)
	}

	t := time.Now()

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("non-2xx response")
	}

	latency <- time.Since(t)

	return nil
}

func listOrderbook(latency chan time.Duration) (*orderbook.OrderBookResponse, error) {
	url := fmt.Sprintf("%sorders/", URL_BASE)

	t := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non-2xx response")
	}

	latency <- time.Since(t)

	out := &orderbook.OrderBookResponse{}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBytes, out)

	return out, err
}
