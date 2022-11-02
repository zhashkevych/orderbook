package orderbook

// Tree Rebalancing (Insert / Delete / Update)
// OrderBooks (pair : OrderBook)

type OrderBook struct {
	Asks *OrderTree
	Bids *OrderTree
}

func NewOrderBook() *OrderBook {
	return &OrderBook{Asks: nil, Bids: nil}
}

func (b *OrderBook) InsertOrder(order *Order) {
	if order.orderType == TYPE_ASK {
		if b.Asks == nil {
			b.Asks = NewOrderTree(order)
		}

		b.Asks.Insert(order)
		return
	}

	if b.Bids == nil {
		b.Bids = NewOrderTree(order)
	}

	b.Bids.Insert(order)
}
