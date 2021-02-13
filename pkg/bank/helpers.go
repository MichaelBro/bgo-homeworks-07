package bank

import "math/rand"

var mccIncludeList = []uint16{4304, 4415, 5122, 5292, 4011, 4112, 5297, 5298, 7911, 7922, 5555, 5556}

var mccCategory = map[string][]uint16{
	"Авиабилеты": {4304, 4415},
	"Аптеки":     {5122, 5292},
	"Железнодорожные билеты": {4011, 4112},
	"Супермаркеты":           {5297, 5298},
	"Развлечения":            {7911, 7922},
	"Финансы":                {5555, 5556},
}

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func generateFakeTransactions(count int) *[]Transaction {
	var transactions []Transaction

	timestamp := 1577836800 // 01.01.2020
	ir := IntRange{1000_00, 100_000_00}
	irMcc := IntRange{0, 12}

	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(int64(timestamp)))
		timestamp += 60 * 60 * 12
		transactions = append(transactions, Transaction{
			Id:        uint64(i),
			UserId:    uint64(irMcc.NextRandom(r)),
			Amount:    int64(ir.NextRandom(r)),
			Timestamp: uint32(timestamp),
			Mcc:       mccIncludeList[irMcc.NextRandom(r)],
		})
	}
	return &transactions
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func contains(slice []uint16, val uint16) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

//TODO: check transaction mcc in group list mcc
func IsMCCInTransaction(mcc string, array []string) (status bool) {
	status = false
	for _, currentMcc := range array {
		if mcc == currentMcc {
			status = true
		}
	}
	return
}
