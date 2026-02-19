package deposit

import "github.com/VehsagriX/bank/pkg/bank/types"

const depositLimit = 50_000

func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= depositLimit {
		card.Balance += amount
	}
	return
}
