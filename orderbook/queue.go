package orderbook

import (
	"container/list"
)

type OrderQueue struct {
	queue *list.List
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{list.New()}
}

func (q *OrderQueue) Push(order *Order) {
	q.queue.PushBack(order)
}

func (q *OrderQueue) Remove(order *Order) {
	for el := q.queue.Front(); el != nil; el = el.Next() {
		if el.Value.(*Order) == order {
			q.queue.Remove(el)
		}
	}
}

func (q *OrderQueue) Update(order *Order) {
	for el := q.queue.Front(); el != nil; el = el.Next() {
		if el.Value.(*Order).ID == order.ID {
			q.queue.InsertAfter(order, el)

			q.queue.Remove(el)
		}
	}
}

func (q *OrderQueue) GetLatest() *Order {
	order := q.queue.Back()

	return order.Value.(*Order)
}

func (q *OrderQueue) GetEarliest() *Order {
	order := q.queue.Front()

	return order.Value.(*Order)
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
