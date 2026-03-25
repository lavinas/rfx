package internal

import (
	"math/rand"
	"time"

	"github.com/thanhpk/randstr"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID                        int64
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	Key1                      string
	EstablishmentCode         int
	EstablishmentNature       int
	EstablishmentMCC          int
	EstablishmentTerminalCode int
	BIN                       int
	AuthorizationCode         string
	TransactionNSU            string
	TransactionDate           time.Time
	TransactionAmount         float64
	TransactionInstallments   int
	TransactionBrand          string
	TransactionProduct        string
	TransactionCapture        string
	RevenueMDRValue           float64
	CostInterchangeValue      float64
	HighSourcePriority        int
	StatusID                  int
	StatusName                string
	StatusCount               int
	PeriodDate                *time.Time
	PeriodClosingID           *int
	TransacID                 string
}

// NewTransaction generates a new  transaction
func NewTransaction() *Transaction {
	return &Transaction{}
}

// GenerateInsert
func (t *Transaction) Insert(db *DB) {
	sql := `insert into transaction.transaction(id, created_at, updated_at, key1, establishment_code, establishment_nature, establishment_mcc, establishment_terminal_code,
	        bin, authorization_code, transaction_nsu, transaction_date, transaction_amount, transaction_installments, transaction_brand,
			transaction_product, transaction_capture, revenue_mdr_value, cost_interchange_value, high_source_priority, status_id, status_name, status_count,
			period_date,period_closing_id,transac_id) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)`
	_, err := db.Exec(sql, t.ID, t.CreatedAt, t.UpdatedAt, t.Key1, t.EstablishmentCode, t.EstablishmentNature, t.EstablishmentMCC, t.EstablishmentTerminalCode,
		t.BIN, t.AuthorizationCode, t.TransactionNSU, t.TransactionDate, t.TransactionAmount, t.TransactionInstallments, t.TransactionBrand,
		t.TransactionProduct, t.TransactionCapture, t.RevenueMDRValue, t.CostInterchangeValue, t.HighSourcePriority, t.StatusID, t.StatusName, t.StatusCount,
		t.PeriodDate, t.PeriodClosingID, t.TransacID)
	if err != nil {
		panic(err)
	}
}

// GetLastID retrieves the last transaction ID from the database
func (t *Transaction) GetLastID(db *DB) int64 {
	var lastID int64
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM transaction.transaction").Scan(&lastID)
	if err != nil {
		panic(err)
	}
	return lastID
}

// Generate random data for the transaction (this is just a placeholder, you can implement it as needed)
func (t *Transaction) GenerateData(id int64, transaction_date time.Time) {
	t.ID = id
	t.TransactionDate = transaction_date
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.Key1 = randstr.String(20)
	t.TransactionDate = transaction_date
	t.AuthorizationCode = randstr.String(6)
	t.TransactionNSU = randstr.String(12)
	t.TransacID = randstr.String(20)
	t.HighSourcePriority = 30
	t.GenerateEstablishmentCode()
	t.GenerateEstablishmentNature()
	t.GenerateEstablishmentMCC()
	t.GenerateEstablishmentTerminalCode()
	t.GenerateBin()
	t.GenerateAmount()
	t.GenerateBrand()
	t.GenerateProduct()
	t.GenerateInstallments()
	t.GenerateCapture()
	t.GenerateStatus()
	t.GenerateMDRValue()
	t.GenerateCostInterchangeValue()
	t.GeneratePeriodDate()
}


// GenerateEstablishmentCode generates a random establishment code between 600067 and 700222
func (t *Transaction) GenerateEstablishmentCode() {
	min := 600067
	max := 700222
	t.EstablishmentCode = rand.Intn(max-min+1) + min
}

// GenerateStatus generates a random status for the transaction
func (t *Transaction) GenerateStatus() {
	rand := rand.Intn(100)
	if rand < 70 {
		t.StatusID = 2 // ready
		t.StatusCount = 1
		t.StatusName = "pronto"
	} else {
		t.StatusID = 1 // pending
		t.StatusCount = 0
		t.StatusName = "pendente"
	}
}

// GenerateEstablishmentNature generates a random establishment nature (0 or 1) with a bias towards 0 (80% chance)
func (t *Transaction) GenerateEstablishmentNature() {
	last := t.EstablishmentCode % 10
	if last < 8 {
		t.EstablishmentNature = 0
	} else {
		t.EstablishmentNature = 1
	}
}

// GenerateEstablishmentMCC generates a random establishment MCC based on the last digit of the establishment code
func (t *Transaction) GenerateEstablishmentMCC() {
	mccs := []int{
		7399, 5199, 5441, 5462, 5811, 1799, 5814, 5541, 8299, 5943,
		5812, 7230, 5732, 5912, 5921, 5942, 7311, 5211, 8099, 5451,
		5651, 5813, 5946, 5499, 7011, 5422, 7216, 7299, 5733, 5411,
	}
	last := t.EstablishmentCode % 30
	t.EstablishmentMCC = mccs[last]
}

// GenerateEstablishmentTerminalCode generates a random establishment terminal code between 1000 and 9999
func (t *Transaction) GenerateEstablishmentTerminalCode() {
	last := rand.Intn(2)
	t.EstablishmentTerminalCode = t.EstablishmentCode*10 + last
}

// GenerateBin generates a random BIN based on the provided values and proportions
func (t *Transaction) GenerateBin() {
	values := []int{
		550209, 439267, 417402, 485464, 516292, 506722, 650032, 589916, 230650, 650570,
		514945, 223656, 555507, 404025, 515467, 650033, 234074, 546997, 223115, 410863,
		417400, 498442, 650489, 498407, 650571, 516362, 477393, 452204, 520132, 522626,
		489389, 476332, 466717, 637529, 544731, 549743, 498453, 559532, 520635, 545960,
		490144, 679002, 550272, 608710, 549289, 521438, 544432, 544584, 512707, 650529,
	}
	proportions := []float64{
		1852, 1087, 285, 279, 278, 277, 241, 175, 155, 132,
		114, 94, 85, 84, 79, 78, 78, 78, 72, 69,
		67, 66, 65, 64, 64, 62, 61, 55, 52, 51,
		50, 50, 48, 48, 47, 46, 45, 45, 43, 41,
		40, 40, 39, 38, 36, 36, 35, 35, 34, 34,
	}
	t.BIN = getProportionValue(values, proportions)
}

// GenerateAmount generates a random transaction amount between 10 and 1000
func (t *Transaction) GenerateAmount() {
	t.TransactionAmount = rand.Float64()*(1000-10) + 10
}

// GenerateBrand generates a random transaction brand based on the provided values and proportions
func (t *Transaction) GenerateBrand() {
	brands := []string{"V", "M", "E"}
	brands_proportions := []float64{0.3, 0.5, 0.2}
	t.TransactionBrand = getProportionValue(brands, brands_proportions)
}

// GenerateProduct generates a random transaction product based on the provided values and proportions
func (t *Transaction) GenerateProduct() {
	products := []string{"DB", "CR"}
	products_proportions := []float64{0.8, 0.2}
	t.TransactionProduct = getProportionValue(products, products_proportions)
}

// GenerateInstallments generates a random number of installments between 1 and 12
func (t *Transaction) GenerateInstallments() {
	if t.TransactionProduct == "DB" {
		t.TransactionInstallments = 1
		return
	}
	inst_values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	inst_proportions := []float64{0.4, 0.15, 0.1, 0.1, 0.05, 0.05, 0.03, 0.02, 0.02, 0.01, 0.01, 0.01}
	t.TransactionInstallments = getProportionValue(inst_values, inst_proportions)
}

// GenerateCapture generates a random transaction capture type based on the provided values and proportions
func (t *Transaction) GenerateCapture() {
	capture_values := []string{"CTC", "CHP"}
	capture_proportions := []float64{2454, 508}
	t.TransactionCapture = getProportionValue(capture_values, capture_proportions)
}

// GenerateMDRValue generates a random revenue MDR value between 0.5 and 5.0
func (t *Transaction) GenerateMDRValue() {
	t.RevenueMDRValue = (rand.Float64()*(5.0-0.5) + 0.5) * t.TransactionAmount / 100
}

// GenerateCostInterchangeValue generates a random cost interchange value between 50% and 80% of the revenue MDR value
func (t *Transaction) GenerateCostInterchangeValue() {
	t.CostInterchangeValue = t.RevenueMDRValue * (rand.Float64()*(0.8-0.5) + 0.5)
}

// ganeratePeriodDate generates a random period date based on the transaction date
func (t *Transaction) GeneratePeriodDate() {
	if t.StatusID == 2 {
		periodDate := t.TransactionDate.AddDate(0, 0, 30)
		t.PeriodDate = &periodDate
		// periodClosingID := 1
		t.PeriodClosingID = nil
	} else {
		t.PeriodDate = nil
		t.PeriodClosingID = nil
	}
}

// getProportionValue returns a random value from the values slice based on the provided proportions
func getProportionValue[T any](values []T, proportions []float64) T {
	randVal := rand.Float64()
	prop_sum := 0.0
	prop_total := 0.0

	for _, p := range proportions {
		prop_total += p
	}

	for i := range proportions {
		prop_sum += proportions[i] / prop_total
		if randVal < prop_sum {
			return values[i]
		}
	}
	return values[len(values)-1]
}
