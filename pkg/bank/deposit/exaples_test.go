package deposit

import (
	"fmt"

	"github.com/VehsagriX/bank/pkg/bank/types"
)

func ExampleDeposit() {
	cardNorm := types.Card{Balance: 20_000, Active: true}
	cardInactive := types.Card{Balance: 20_000, Active: false}
	Deposit(&cardNorm, 1_000)
	Deposit(&cardNorm, 50_001)
	Deposit(&cardInactive, 1_000)
	fmt.Println(cardNorm.Balance, cardNorm.Balance, cardInactive.Balance)
	//Output: 21000 21000 20000

}
