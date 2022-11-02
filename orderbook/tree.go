package orderbook

import "fmt"

// OrderTree implements BST data structure
type OrderTree struct {
	orders *OrderQueue
	price  int

	left  *OrderTree
	right *OrderTree
}

func NewOrderTree(order *Order) *OrderTree {
	orders := NewOrderQueue()
	orders.Push(order)

	return &OrderTree{
		orders: orders,
		price:  order.Price,
		left:   nil,
		right:  nil,
	}
}

func (t *OrderTree) Insert(order *Order) {
	if order.Price == t.price {
		t.orders.Push(order)

		return
	}

	if order.Price <= t.price {
		if t.left == nil {
			t.left = NewOrderTree(order)
		}

		t.left.Insert(order)

		return
	}

	if t.right == nil {
		t.right = NewOrderTree(order)
	}

	t.right.Insert(order)
}

func (t *OrderTree) Remove(order *Order) {
	if order.Price == t.price {
		t.orders.Remove(order)

		return
	}

	if order.Price <= t.price {
		if t.left != nil {
			t.left.Remove(order)
		}

		return
	}

	if t.right != nil {
		t.right.Remove(order)
	}
}

func (t *OrderTree) Update(order *Order) {
	if order.Price == t.price {
		t.orders.Update(order)

		return
	}

	if order.Price <= t.price {
		if t.left != nil {
			t.left.Update(order)
		}

		return
	}

	if t.right != nil {
		t.right.Update(order)
	}
}

func (t *OrderTree) FindMinPrice() int {
	if t.left == nil {
		return t.price
	}

	return t.left.FindMinPrice()
}

func (t *OrderTree) FindMaxPrice() int {
	if t.right == nil {
		return t.price
	}

	return t.left.FindMaxPrice()
}

func (t *OrderTree) Find(price int) (*OrderTree, bool) {
	if t == nil {
		return &OrderTree{}, false
	}

	switch {
	case price == t.price:
		return t, true
	case price < t.price:
		return t.left.Find(price)
	default:
		return t.right.Find(price)
	}
}

func (t *OrderTree) PrintInorder() {
	if t == nil {
		return
	}

	t.left.PrintInorder()

	fmt.Printf("Price: %d | Number of orders: %d\n", t.price, t.orders.Len())
	orders := t.orders.GetAll()
	for i, order := range orders {
		fmt.Printf("\t#%d. Amount: %d\n", i+1, order.Amount)
	}
	fmt.Println()

	t.right.PrintInorder()
}
