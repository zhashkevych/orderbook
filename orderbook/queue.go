package orderbook

import "container/list"

type OrderQueue struct {
	queue *list.List

	price int
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{list.New(), 0}
}

func (q *OrderQueue) Push(order *Order) {
	q.queue.PushFront(order)

	q.price = order.price
}

func (q *OrderQueue) GetLatest() *Order {
	order := q.queue.Front()

	return order.Value.(*Order)
}

func (q *OrderQueue) GetPrice() int {
	return q.price
}

func (q *OrderQueue) Len() int {
	return q.queue.Len()
}

func (q *OrderQueue) GetAll() []*Order {
	out := make([]*Order, q.queue.Len())
	index := 0

	for order := q.queue.Front(); order != nil; order = order.Next() {
		out[index] = order.Value.(*Order)
		index++
	}

	return out
}
