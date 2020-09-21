package stats

import (
	"github.com/nazurdinov95/bank/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	for _, payment := range payments {
		avg += payment.Amount
	}

	return avg / types.Money (len(payments))
}