package orderbook

type Orderbook struct {
	Ask      *Heap
	Bid      *Heap
	TotalAsk uint64
	TotalBid uint64
}

func New() *Orderbook {

	ask := &Heap{comp: less, elems: make([]*Order, 0)}
	bid := &Heap{comp: greater, elems: make([]*Order, 0)}
	return &Orderbook{Ask: ask, Bid: bid, TotalAsk: 0, TotalBid: 0}
}

type Condition struct {
	total, rtotal *uint64
	heap, rheap   *Heap
	price         uint64
}

func CreateCond(orderbook *Orderbook, order *Order) *Condition {
	var C *Condition
	if order.Side.String() == "ASK" {
		C = &Condition{&orderbook.TotalAsk, &orderbook.TotalBid,
			orderbook.Bid, orderbook.Ask, 0}

	} else if order.Side.String() == "BID" {
		C = &Condition{&orderbook.TotalBid, &orderbook.TotalAsk,
			orderbook.Ask, orderbook.Bid, 1e18}
	}
	return C
}

func (orderbook *Orderbook) Process(order *Order, cond *Condition) ([]*Trade, *Order) {
	var trade []*Trade

	for order.Volume != 0 && !cond.heap.IsEmpty() && cond.heap.comp(cond.heap.Top(), order) {
		if order.Volume >= cond.heap.Top().Volume {
			order.Volume -= cond.heap.Top().Volume
			*(cond.rtotal) -= cond.heap.Top().Volume
			last := cond.heap.Pop()

			trade = append(trade, &Trade{Volume: last.Volume, Price: last.Price})

		} else {

			cond.heap.Top().Volume -= order.Volume
			*(cond.rtotal) -= order.Volume
			trade = append(trade, &Trade{Volume: order.Volume, Price: cond.heap.Top().Price})
			order.Volume = 0
		}
	}

	if order.Volume != 0 {
		if order.Kind.String() == "MARKET" {
			return trade, order
		}

		cond.rheap.Insert(order)
		*(cond.total) += order.Volume
	}

	return trade, nil
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	var cond *Condition
	cond = CreateCond(orderbook, order)

	if order.Kind.String() == "MARKET" {
		order.Price = cond.price
	}

	return orderbook.Process(order, cond)
}

//Heap implemention

func less(a, b *Order) bool {

	if (*a).Price != (*b).Price {
		return (*a).Price < (*b).Price
	}

	return (*a).ID < (*b).ID

}
func greater(a, b *Order) bool {

	if (*a).Price != (*b).Price {
		return (*a).Price > (*b).Price
	}

	return (*a).ID < (*b).ID

}

type Heap struct {
	comp  func(a, b *Order) bool
	elems []*Order
}

func (heap *Heap) Heapify(c_ind int) {

	l_ind := c_ind*2 + 1
	r_ind := c_ind*2 + 2

	comp := heap.comp
	size := len(heap.elems)

	c_el := &(heap.elems[c_ind])

	n_el := c_el
	n_ind := c_ind

	if l_ind < size {
		l_el := &(heap.elems[l_ind])
		if comp(*l_el, *n_el) {
			n_el = l_el
			n_ind = l_ind
		}
	}
	if r_ind < size {
		r_el := &(heap.elems[r_ind])
		if comp(*r_el, *n_el) {
			n_el = r_el
			n_ind = r_ind
		}
	}

	if n_el == c_el {
		return
	}

	*n_el, *c_el = *c_el, *n_el

	heap.Heapify(n_ind)

}

func (heap *Heap) Up(c_ind int) {

	if c_ind == 0 {
		return
	}

	comp := heap.comp

	n_ind := (c_ind - 1) / 2

	c_el := &(heap.elems[c_ind])
	n_el := &(heap.elems[n_ind])

	if comp(*c_el, *n_el) {
		*n_el, *c_el = *c_el, *n_el
		heap.Up(n_ind)
	}

}

func (heap *Heap) Insert(order *Order) {

	heap.elems = append(heap.elems, order)
	heap.Up(len(heap.elems) - 1)

}

func (heap *Heap) IsEmpty() bool {
	if len(heap.elems) == 0 {
		return true
	}
	return false
}

func (heap *Heap) Top() *Order {
	return heap.elems[0]
}

func (heap *Heap) Pop() *Order {

	head := heap.Top()

	heap.elems[0] = heap.elems[len(heap.elems)-1]
	heap.elems = heap.elems[:len(heap.elems)-1]

	if len(heap.elems) != 0 {
		heap.Heapify(0)
	}

	return head

}
