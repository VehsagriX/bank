package card

import "github.com/VehsagriX/bank/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Name:     name,
		Active:   true,
		Color:    color,
	}
	return card
}
