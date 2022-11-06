package orderbook

import (
	"time"

	"github.com/google/uuid"
)

type OrderType int

const (
	TYPE_ASK OrderType = iota
	TYPE_BID
)

type Order struct {
	ID        uuid.UUID `json:"id"`
	Price     int       `json:"price"`
	Amount    int       `json:"amount"`
	OrderType OrderType `json:"order_type"`
	CreatedAt time.Time
}

func NewOrder(price, amount int, orderType OrderType) *Order {
	return &Order{uuid.New(), price, amount, orderType, time.Now()}
}
