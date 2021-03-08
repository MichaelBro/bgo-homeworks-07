package bank

import "math/rand"

var mccIncludeList = []Mcc{
	{
		Code:     4304,
		Category: "Авиабилеты",
	},
	{
		Code:     5122,
		Category: "Аптеки",
	},
	{
		Code:     4304,
		Category: "Железнодорожные билеты",
	},
	{
		Code:     5298,
		Category: "Супермаркеты",
	},
	{
		Code:     7911,
		Category: "Развлечения",
	},
	{
		Code:     5555,
		Category: "Финансы",
	},
}

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func GenerateRandomTransactions(count int) []Transaction {
	var transactions []Transaction

	timestamp := 1577836800 // 01.01.2020
	ir := IntRange{10, 5000}
	irMcc := IntRange{1, len(mccIncludeList) - 1}

	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(int64(timestamp)))
		timestamp += 60 * 60 * 24
		transactions = append(transactions, Transaction{
			Id:        uint64(i),
			UserId:    uint64(irMcc.NextRandom(r)),
			Amount:    int64(ir.NextRandom(r)),
			Timestamp: uint32(timestamp),
			Mcc:       mccIncludeList[irMcc.NextRandom(r)],
		})
	}
	return transactions
}
