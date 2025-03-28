package main

import (
	"time"
)

type OrderSide i8
type OrderType i8

const BID OrderSide = 0
const ASK OrderSide = 1
const LIMIT OrderType = 0
const MARKET OrderType = 1

type Order struct {
	id         u64
	otype      OrderType
	side       OrderSide
	size       i32
	price      f32
	created    time.Time
	filled_pct f32
}

func (order *Order) Fill(size i32) (i32, error) {
	// would execute the ownership transfer
	if size > order.size {
		return 0, &BaseError{message: "Filling size exceeds order size"}
	} else {
		order.filled_pct = f32(size) / f32(order.size)
		order.size = order.size - size
		return size, nil
	}
}
