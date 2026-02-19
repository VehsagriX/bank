package transfer

import "github.com/VehsagriX/bank/pkg/bank/types"

// Bonus расчитываает по amount начисляемый при переводе бонус.
func Bonus(amount types.Money) types.Money {
	const bonusPrecent = 0.0050
	bonus := types.Money(float64(amount) * bonusPrecent)
	return bonus
}

// Total расчитывает по amount итоговую сумму для зачисления с учетом бонуса
func Total(amount types.Money) types.Money {
	total := amount + Bonus(amount)
	return total
}
