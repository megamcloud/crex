package gotrader

import (
	. "github.com/coinrust/gotrader/models"
)

type Broker interface {
	GetAccountSummary(currency string) (result AccountSummary, err error)
	GetOrderBook(symbol string, depth int) (result OrderBook, err error)
	PlaceOrder(symbol string, direction Direction, orderType OrderType, price float64, amount float64,
		postOnly bool, reduceOnly bool) (result Order, err error)
	GetOpenOrders(symbol string) (result []Order, err error)
	GetOrder(symbol string, id uint64) (result Order, err error)
	GetPosition(symbol string) (result Position, err error)
}