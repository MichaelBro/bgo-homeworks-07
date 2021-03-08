package main

import (
	"bgo-homeworks-07/pkg/bank"
	"bgo-homeworks-07/pkg/processing"
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	transactions := bank.GenerateRandomTransactions(100000)
	userId := uint64(5)
	g := 5

	SumTransactionsByCategory := processing.SumTransactionsByCategory(transactions, userId)
	fmt.Println("SumTransactionsByCategory:              ", SumTransactionsByCategory)

	SumCategoryTransactionsMutex := processing.SumCategoryTransactionsMutex(transactions, userId, g)
	fmt.Println("SumCategoryTransactionsMutex:           ", SumCategoryTransactionsMutex)

	SumCategoryTransactionsChanel := processing.SumCategoryTransactionsChanel(transactions, userId, g)
	fmt.Println("SumCategoryTransactionsChanel:          ", SumCategoryTransactionsChanel)

	SumCategoryTransactionsMutexStandalone := processing.SumCategoryTransactionsMutexStandalone(transactions, userId, g)
	fmt.Println("SumCategoryTransactionsMutexStandalone: ", SumCategoryTransactionsMutexStandalone)
}
