package withdraw

import (
	"fmt"

	"github.com/VehsagriX/bank/v2/pkg/bank/types"
)

func ExampleWithDraw() {
	card := types.Card{Balance: 20_000_00, Active: true}
	WithDraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	//Output: 1000000

}
