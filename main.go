package main

import (
	"fmt"

	"github.com/zhashkevych/orderbook/orderbook"
)

func main() {
	orderBook := orderbook.NewOrderBook()

	// BIDS
	orderBook.InsertOrder(orderbook.NewOrder(30, 10, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(30, 8, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(30, 15, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(31, 1, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(31, 145, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(32, 2, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(30, 18, orderbook.TYPE_BID))
	orderBook.InsertOrder(orderbook.NewOrder(31, 4, orderbook.TYPE_BID))

	// ASKS
	orderBook.InsertOrder(orderbook.NewOrder(31, 4, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(32, 5, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(32, 5, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(32, 8, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(33, 8, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(33, 198, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(36, 1, orderbook.TYPE_ASK))

	// fmt.Printf("%+v", orderBook.BuyOrders.order)
	fmt.Println("| BIDS |")
	orderBook.Bids.PrintInorder()

	fmt.Println()

	fmt.Println("| ASKS |")
	orderBook.Asks.PrintInorder()
}
