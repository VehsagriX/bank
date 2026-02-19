package total

import (
	"fmt"

	"github.com/VehsagriX/bank/pkg/bank/types"
)

func ExampleTotal() {
	// Использование нижнего подчеркивания в числах (10_000_00)
	// разрешено в Go для читаемости.
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}

	fmt.Println(Total(cards))
	// Output: 2000000
}
