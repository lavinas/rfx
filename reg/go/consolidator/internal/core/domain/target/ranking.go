package target_domain

import (
	"fmt"
	"time"
	"sort"

	source_domain "consolidator/internal/core/domain/source"
)

// Ranking represents the ranking of establishments based on transaction data.
type Ranking struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	EstablishmentCode   int64     `gorm:"column:establishment_code"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         int       `gorm:"column:segment_code"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int       `gorm:"column:transaction_quantity"`
	AvgMccFee           float64   `gorm:"column:avg_mcc_fee"`
}

// NewRanking creates a new instance of Ranking.
func NewRanking() *Ranking {
	return &Ranking{}
}

// TableName returns the name of the database table for the Ranking model.
func (i *Ranking) TableName() string {
	return "cadoc_6334_v2.ranking"
}

// GetFromTransaction returns a Ranking instance populated with data from a given transaction.
func (i *Ranking) GetFromTransaction(transaction *source_domain.Transaction) *Ranking {
	return &Ranking{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		EstablishmentCode:   transaction.GetEstablishmentCode(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
		AvgMccFee:           transaction.GetRevenueMDRValueRate(),
	}
}

// GetKey generates a unique key for the Ranking struct based on its fields.
func (i *Ranking) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// GetFromTransactions processes a slice of transactions and returns a map of Ranking instances keyed by their unique keys.
func (i *Ranking) AddTransactions(transactions []*source_domain.Transaction, items map[string]*Ranking) {
	for _, t := range transactions {
		ranking := i.GetFromTransaction(t)
		key := ranking.GetKey()
		if existing, exists := items[key]; exists {
			existing.TransactionAmount += ranking.TransactionAmount
			existing.TransactionQuantity += ranking.TransactionQuantity
			delta := ranking.AvgMccFee - existing.AvgMccFee
			existing.AvgMccFee += delta / float64(existing.TransactionQuantity)
		} else {
			items[key] = ranking
		}
	}
}

// FilterRanking filters the ranking data to include only the top and bottom establishments based on transaction amount
func FilterRanking(items map[string]*Ranking) map[string]*Ranking {
	filtered := make(map[string]*Ranking)

	// Group establishments by segment code and sum transaction amounts
	segmentEstablishments := make(map[int]map[int64]float64)
	for _, ranking := range items {
		if _, exists := segmentEstablishments[ranking.SegmentCode]; !exists {
			segmentEstablishments[ranking.SegmentCode] = make(map[int64]float64)
		}
		segmentEstablishments[ranking.SegmentCode][ranking.EstablishmentCode] += ranking.TransactionAmount
	}

	// for each segment code, filter the 15 first 
	for segmentCode := range segmentEstablishments {
		// sort establishments by transaction amount
		type establishment struct {
			Code   int64
			Amount float64
		}
		var establishments []establishment
		for code, amount := range segmentEstablishments[segmentCode] {
			establishments = append(establishments, establishment{Code: code, Amount: amount})
		}
		sort.Slice(establishments, func(i, j int) bool {
			return establishments[i].Amount > establishments[j].Amount
		})

		// get top 15 establishments
		topCount := 15
		if len(establishments) < topCount {
			topCount = len(establishments)
		}
		for i := 0; i < topCount; i++ {
			code := establishments[i].Code
			for key, ranking := range items {
				if ranking.SegmentCode == segmentCode && ranking.EstablishmentCode == code {
					filtered[key] = ranking
				}
			}
		}
		
		// get last 200 establishments
		filtered_200 := make(map[string]*Ranking) 
		topCount = 200
		if len(establishments) < topCount {
			topCount = len(establishments)
		}
		for i := len(establishments) - topCount; i < len(establishments); i++ {
			code := establishments[i].Code
			for key, ranking := range items {
				if ranking.SegmentCode == segmentCode && ranking.EstablishmentCode == code {
					filtered_200[key] = ranking
				}
			}
		}

		// consolidate last 200 establishments and add to filtered
		filtered_200_consolidate := consolidateRanking(filtered_200)
		for key, ranking := range filtered_200_consolidate 	{
			filtered[key] = ranking
		}

	}
	return filtered

}

// Consolidate Ranking 
func consolidateRanking(ranking map[string]*Ranking) map[string]*Ranking {
	consRanking := make(map[string]*Ranking)
	for _, r := range ranking {
		NewRanking := &Ranking{
			Year:                r.Year,
			Quarter:             r.Quarter,
			EstablishmentCode:   -1,
			Function:            r.Function,
			Brand:               r.Brand,
			CaptureMode:         r.CaptureMode,
			Installments:        r.Installments,
			SegmentCode:         r.SegmentCode,
			TransactionAmount:   r.TransactionAmount,
			TransactionQuantity: r.TransactionQuantity,
			AvgMccFee:           r.AvgMccFee,
		}
		key := NewRanking.GetKey()
		if existing, exists := consRanking[key]; exists {
			existing.TransactionAmount += r.TransactionAmount
			existing.TransactionQuantity += r.TransactionQuantity
			delta := r.AvgMccFee - existing.AvgMccFee
			existing.AvgMccFee += delta / float64(existing.TransactionQuantity)
		} else {
			consRanking[key] = NewRanking
		}
	}
	return consRanking
}
