package main

import (
	"bgo-homeworks-07/pkg/bank"
	"bgo-homeworks-07/pkg/processing"
	"fmt"
)

func main() {
	transactions := bank.GenerateFakeTransactions(100000)

	g := 4
	sumByCategoriesUser1, err := processing.SumCategoryTransactionsMutex(transactions, uint64(1), g)
	if err != nil {
		fmt.Println(err)
	}
	sumByCategoriesUser2, err := processing.SumCategoryTransactionsMutex(transactions, uint64(2), g)
	if err != nil {
		fmt.Println(err)
	}
	sumByCategoriesUser3, err := processing.SumCategoryTransactionsMutex(transactions, uint64(3), g)
	if err != nil {
		fmt.Println(err)
	}
	sumByCategoriesUser4, err := processing.SumCategoryTransactionsMutex(transactions, uint64(4), g)
	if err != nil {
		fmt.Println(err)
	}
	sumByCategoriesUser5, err := processing.SumCategoryTransactionsMutex(transactions, uint64(4), g)
	if err != nil {
		fmt.Println(err)
	}
	sumByCategoriesUser6, err := processing.SumCategoryTransactionsMutex(transactions, uint64(4), g)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sumByCategoriesUser1: ", sumByCategoriesUser1)
	fmt.Println("sumByCategoriesUser2: ", sumByCategoriesUser2)
	fmt.Println("sumByCategoriesUser3: ", sumByCategoriesUser3)
	fmt.Println("sumByCategoriesUser4: ", sumByCategoriesUser4)
	fmt.Println("sumByCategoriesUser5: ", sumByCategoriesUser5)
	fmt.Println("sumByCategoriesUser6: ", sumByCategoriesUser6)
}
