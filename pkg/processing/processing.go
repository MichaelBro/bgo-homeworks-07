package processing

import (
	"bgo-homeworks-07/pkg/bank"
	"errors"
	"sync"
)

var (
	ErrNoTransactions = errors.New("empty transactions")
)

func SumTransactionsByCategory(transactions []bank.Transaction, userId uint64) map[string]int64 {
	if transactions == nil {
		return nil
	}

	m := make(map[string]int64)
	for _, transaction := range transactions {
		if transaction.UserId == userId {
			m[transaction.Mcc.Category] += transaction.Amount
		}
	}
	return m
}

func SumCategoryTransactionsMutex(transactions []bank.Transaction, userId uint64, goroutines int) map[string]int64 {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	wg := sync.WaitGroup{}
	wg.Add(goroutines)
	mu := sync.Mutex{}
	result := make(map[string]int64)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			m := SumTransactionsByCategory(part, userId)

			mu.Lock()
			for key, value := range m {
				result[key] += value
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}

func SumCategoryTransactionsChanel(transactions []bank.Transaction, userId uint64, goroutines int) map[string]int64 {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	result := make(map[string]int64)
	ch := make(chan map[string]int64)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[string]int64) {
			ch <- SumTransactionsByCategory(part, userId)
		}(ch)
	}

	finished := 0
	for sum := range ch {
		for k, v := range sum {
			result[k] += v

		}
		finished++
		if finished == goroutines {
			break
		}
	}

	return result
}

func SumCategoryTransactionsMutexStandalone(transactions []bank.Transaction, userId uint64, goroutines int) map[string]int64 {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		wg.Add(1)

		go func() {
			for _, transaction := range part {
				if transaction.UserId == userId {
					mu.Lock()
					result[transaction.Mcc.Category] += transaction.Amount
					mu.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}
