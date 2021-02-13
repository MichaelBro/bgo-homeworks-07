package processing

import (
	"bgo-homeworks-07/pkg/bank"
	"errors"
	"sync"
)

var (
	ErrNoTransactions = errors.New("no user transactions")
)

func SumTransactions(transactions []bank.Transaction, userId uint64) map[string]int64 {
	m := make(map[string]int64)
	for _, transaction := range transactions {
		if transaction.UserId == userId {
			m[transaction.Mcc.Category] += transaction.Amount
		}
	}
	return m
}

func SumCategoryTransactionsMutex(transactions []bank.Transaction, userId uint64, goroutines int) (map[string]int64, error) {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	m := make(map[string]int64)

	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			mapSum := SumTransactions(part, userId)

			mu.Lock()
			for key, i := range mapSum {
				m[key] += i

			}
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()

	return m, nil
}
