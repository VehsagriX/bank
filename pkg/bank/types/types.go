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

type Status string

const (
	StatusOK         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
