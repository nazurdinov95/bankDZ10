package stats

import (
	"github.com/nazurdinov95/bank/pkg/bank/types"
	"fmt"
)


func ExampleAvg()  {
	payments := []types.Payment{
	{
		ID: 1,
		Amount: 100,
		Category: "авто",
	},
	{
		ID: 1,
		Amount: 200,
		Category: "аптеки",
	},
}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 150
}

func ExampleTotalInCategory()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 100,
			Category: "авто",
		},
		{
			ID: 1,
			Amount: 200,
			Category: "аптеки",
		},
	}
		result := TotalInCategory(payments, "авто")
		fmt.Println(result)
		//Output: 100
}
