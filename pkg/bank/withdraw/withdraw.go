package withdraw

import (
	"github.com/VehsagriX/bank/v2/pkg/bank/types"
)

const withDrawLimit = 20_000_00

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
