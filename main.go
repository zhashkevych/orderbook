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
	order := orderbook.NewOrder(31, 4, orderbook.TYPE_ASK)
	orderBook.InsertOrder(order)

	orderBook.InsertOrder(orderbook.NewOrder(32, 5, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(32, 5, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(32, 8, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(33, 8, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(33, 198, orderbook.TYPE_ASK))
	orderBook.InsertOrder(orderbook.NewOrder(36, 1, orderbook.TYPE_ASK))

	fmt.Println("| ASKS BEFORE REMOVE |")
	orderBook.Asks.PrintInorder()

	orderBook.RemoveOrder(order)

	fmt.Println("| ASKS AFTER REMOVE |")
	orderBook.Asks.PrintInorder()

	orderBook.InsertOrder(orderbook.NewOrder(31, 11, orderbook.TYPE_ASK))
	orderBook.InsertOrder(order)
	orderBook.InsertOrder(orderbook.NewOrder(31, 22, orderbook.TYPE_ASK))

	fmt.Println("| ASKS AFTER INSERT |")
	orderBook.Asks.PrintInorder()

	order.Amount = 15
	orderBook.UpdateOrder(order)

	fmt.Println("| ASKS AFTER UPDATE |")
	orderBook.Asks.PrintInorder()
}
