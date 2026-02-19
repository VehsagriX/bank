package total

import "github.com/VehsagriX/bank/pkg/bank/types"

func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if !card.Active || card.Balance <= 0 {
			continue
		}

		total += card.Balance
	}
	return total
}
