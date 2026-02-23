package withdraw

import "github.com/VehsagriX/bank/pkg/bank/types"

const withDrawLimit = 20_000_00

// Снятие данных с карты
func WithDraw(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}
	if amount > withDrawLimit {
		return
	}
	if card.Active == false {
		return
	}
	if card.Balance < amount {
		return
	}

	card.Balance -= amount

}
