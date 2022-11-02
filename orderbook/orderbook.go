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
	if order.OrderType == TYPE_ASK {
		if b.Asks == nil {
			b.Asks = NewOrderTree(order)

			return
		}

		b.Asks.Insert(order)

		return
	}

	if b.Bids == nil {
		b.Bids = NewOrderTree(order)

		return
	}

	b.Bids.Insert(order)
}

func (b *OrderBook) RemoveOrder(order *Order) {
	if order.OrderType == TYPE_ASK {
		if b.Asks != nil {
			b.Asks.Remove(order)
		}
	}

	if b.Bids != nil {
		b.Bids.Remove(order)
	}
}

func (b *OrderBook) UpdateOrder(order *Order) {
	if order.OrderType == TYPE_ASK {
		if b.Asks != nil {
			b.Asks.Update(order)
		}
	}

	if b.Bids != nil {
		b.Bids.Update(order)
	}
}
