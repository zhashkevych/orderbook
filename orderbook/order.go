package orderbook

import "time"

type OrderType int

const (
	TYPE_ASK OrderType = iota
	TYPE_BID
)

type Order struct {
	price     int
	amount    int
	orderType OrderType
	createdAt time.Time
}

func NewOrder(price, amount int, orderType OrderType) *Order {
	return &Order{price, amount, orderType, time.Now()}
}
