package payment

import (
	"github.com/VehsagriX/bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	maxPayment := payments[0]
	for _, payment := range payments[1:] {
		if maxPayment.Amount < payment.Amount {
			maxPayment = payment
		}
	}

	return maxPayment
}
