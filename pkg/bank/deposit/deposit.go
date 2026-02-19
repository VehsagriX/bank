package deposit

import "github.com/VehsagriX/bank/pkg/bank/types"

func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= 50_000 {
		card.Balance += amount
	}
	return
}
