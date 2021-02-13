package bank

import "sort"

type Transaction struct {
	Id        uint64
	UserId    uint64
	Amount    int64
	Timestamp uint32
	Mcc       Mcc
}

type List struct {
	Transactions []Transaction
}

type Mcc struct {
	Code     uint16
	Category string
}

func (l *List) Sort() []Transaction {
	sort.SliceStable(l.Transactions, func(i, j int) bool {
		return l.Transactions[i].Amount > l.Transactions[j].Amount
	})
	return l.Transactions
}
