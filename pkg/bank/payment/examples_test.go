package payment

import (
	"fmt"

	"github.com/VehsagriX/bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     11,
			Amount: 1000,
		},
		{
			ID:     12,
			Amount: 21,
		},
		{
			ID:     99,
			Amount: 11212,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 99
}
