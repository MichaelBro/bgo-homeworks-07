package processing

import (
	"bgo-homeworks-07/pkg/bank"
	"reflect"
	"testing"
)

const testUserId uint64 = 5

var staticTransactions = GenerateStaticTransactions(100, 100_000, 1)

func GenerateStaticTransactions(users, transactionsPerUser, transactionAmount int) []bank.Transaction {
	var timestamp = 1577836800 // 01.01.2020
	transactions := make([]bank.Transaction, users*transactionsPerUser)
	for index := range transactions {
		timestamp += 60 * 60 * 24

		switch index % 100 {
		case 0:
			// Например, каждая 100-ая транзакция в банке от нашего юзера в категории Развлечения
			transactions[index] = bank.Transaction{
				Id:        uint64(index),
				UserId:    testUserId,
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Mcc: bank.Mcc{
					Code:     7911,
					Category: "Развлечения",
				},
			}
		case 20:
			// Например, каждая 120-ая транзакция в банке от нашего юзера в категории Железнодорожные билеты
			transactions[index] = bank.Transaction{
				Id:        uint64(index),
				UserId:    testUserId,
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Mcc: bank.Mcc{
					Code:     4304,
					Category: "Железнодорожные билеты",
				},
			}
		default:
			// Транзакции других юзеров, нужны для "общей" массы
			transactions[index] = bank.Transaction{
				Id:        uint64(index),
				UserId:    uint64(index + 5),
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Mcc: bank.Mcc{
					Code:     5298,
					Category: "Супермаркеты",
				},
			}
		}
	}
	return transactions
}

type args struct {
	transactions []bank.Transaction
	userId       uint64
	goroutines   int
}

var tests = []struct {
	name    string
	args    args
	want    map[string]int64
	wantErr error
}{
	{
		name: "#1 nil transactions",
		args: args{
			transactions: nil,
			userId:       testUserId,
			goroutines:   4,
		},
		want: nil,
	},
	{
		name: "#2 one transaction belongs not needed user",
		args: args{
			transactions: []bank.Transaction{
				{
					Id:        uint64(3),
					UserId:    uint64(6),
					Amount:    int64(100),
					Timestamp: uint32(1577836800),
					Mcc: bank.Mcc{
						Code:     5298,
						Category: "Супермаркеты",
					},
				},
			},
			userId:     testUserId,
			goroutines: 4,
		},
		want: map[string]int64{},
	},
	{
		name: "#3 one transaction belongs needed user",
		args: args{
			transactions: []bank.Transaction{
				{
					Id:        uint64(3),
					UserId:    testUserId,
					Amount:    int64(100),
					Timestamp: uint32(1577836800),
					Mcc: bank.Mcc{
						Code:     5298,
						Category: "Супермаркеты",
					},
				},
			},
			userId:     testUserId,
			goroutines: 4,
		},
		want: map[string]int64{"Супермаркеты": 100},
	},
	{
		name: "#4 three transaction, has one belongs needed user",
		args: args{
			transactions: []bank.Transaction{
				{
					Id:        uint64(1),
					UserId:    uint64(1),
					Amount:    int64(100),
					Timestamp: uint32(1577836800),
					Mcc: bank.Mcc{
						Code:     5298,
						Category: "Супермаркеты",
					},
				},
				{
					Id:        uint64(2),
					UserId:    testUserId,
					Amount:    int64(500),
					Timestamp: uint32(1577836800),
					Mcc: bank.Mcc{
						Code:     7911,
						Category: "Развлечения",
					},
				},
				{
					Id:        uint64(3),
					UserId:    uint64(1),
					Amount:    int64(100),
					Timestamp: uint32(1577836800),
					Mcc: bank.Mcc{
						Code:     5298,
						Category: "Супермаркеты",
					},
				},
			},
			userId:     testUserId,
			goroutines: 4,
		},
		want: map[string]int64{"Развлечения": 500},
	},
	{
		name: "#5 sum one million transactions per user",
		args: args{
			transactions: staticTransactions,
			userId:       testUserId,
			goroutines:   4,
		},
		want: map[string]int64{"Железнодорожные билеты": 100000, "Развлечения": 100000},
	},
}

func TestSumCategoryTransactionsChanel(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumCategoryTransactionsChanel(tt.args.transactions, tt.args.userId, tt.args.goroutines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsChanel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCategoryTransactionsMutex(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumCategoryTransactionsMutex(tt.args.transactions, tt.args.userId, tt.args.goroutines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsMutex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCategoryTransactionsMutexStandalone(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumCategoryTransactionsMutexStandalone(tt.args.transactions, tt.args.userId, tt.args.goroutines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsMutexStandalone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTransactionsByCategory(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTransactionsByCategory(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumTransactionsByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

var wantForBenchmark = map[string]int64{"Железнодорожные билеты": 100000, "Развлечения": 100000}

func BenchmarkSumTransactionsByCategory(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumTransactionsByCategory(staticTransactions, testUserId)
		b.StopTimer()
		if !reflect.DeepEqual(result, wantForBenchmark) {
			b.Fatalf("invalid result, got %v, want %v", result, wantForBenchmark)
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsChanel(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsChanel(staticTransactions, testUserId, 4)
		b.StopTimer()
		if !reflect.DeepEqual(result, wantForBenchmark) {
			b.Fatalf("invalid result, got %v, want %v", result, wantForBenchmark)
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsMutex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsMutex(staticTransactions, testUserId, 4)
		b.StopTimer()
		if !reflect.DeepEqual(result, wantForBenchmark) {
			b.Fatalf("invalid result, got %v, want %v", result, wantForBenchmark)
		}
		b.StartTimer()
	}
}

func BenchmarkSumCategoryTransactionsMutexStandalone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumCategoryTransactionsMutexStandalone(staticTransactions, testUserId, 4)
		b.StopTimer()
		if !reflect.DeepEqual(result, wantForBenchmark) {
			b.Fatalf("invalid result, got %v, want %v", result, wantForBenchmark)
		}
		b.StartTimer()
	}
}
