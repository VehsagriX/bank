package types

// Money представляяет собой денежную сумму
type Money int64

// PAN представляет номер карты
type PAN string

// Currency представляет код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Payment struct {
	ID     int
	Amount Money
}

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
