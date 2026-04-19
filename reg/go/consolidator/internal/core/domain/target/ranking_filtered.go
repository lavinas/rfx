package target_domain

import (
	"fmt"
	"sort"
	"time"
)

const (
	// TopCount defines the number of top establishments to include in the filtered ranking
	topCount = 15
	// BottomCount defines the number of bottom establishments to include in the filtered ranking
	bottomCount = 200
)

// RankingFiltered represents the filtered ranking of establishments based on transaction data.
type RankingFiltered struct {
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

// establishment represents the data structure for establishments which will be used for filtering the ranking data
type establishment struct {
	Code   int64
	Amount float64
}

// NewRankingFiltered creates a new instance of RankingFiltered.
func NewRankingFiltered() *RankingFiltered {
	return &RankingFiltered{}
}

// GetKey generates a unique key for the Ranking struct based on its fields.
func (i *RankingFiltered) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// TableName returns the name of the database table for the RankingFiltered model.
func (i *RankingFiltered) TableName() string {
	return "cadoc_6334_v2.ranking_filtered"
}

// FilterRanking filters the ranking data to include only the top and bottom establishments based on transaction amount
func FilterRanking(items map[string]*Ranking) map[string]*RankingFiltered {
	filtered := make(map[string]*RankingFiltered)

	// Group establishments by segment code and sum transaction amounts
	segmentEstablishments := make(map[int]map[int64]float64)
	for _, ranking := range items {
		if _, exists := segmentEstablishments[ranking.SegmentCode]; !exists {
			segmentEstablishments[ranking.SegmentCode] = make(map[int64]float64)
		}
		segmentEstablishments[ranking.SegmentCode][ranking.EstablishmentCode] += ranking.TransactionAmount
	}

	// for each segment code, filter top and bottom establishments
	for segmentCode := range segmentEstablishments {
		// sort establishments by transaction amount
		establishments := getRankingSortedEstablishments(segmentCode, segmentEstablishments)

		// get top establishments
		filtered = filterTopRanking(items, segmentCode, establishments)

		// get bottom establishments
		bottom := filterBottomRanking(items, segmentCode, establishments)
		// consolidate bottom establishments
		bottom = consolidateBottomRanking(bottom)

		// add bottom establishments to filtered
		for key, bt := range bottom {
			filtered[key] = bt
		}

	}

	return filtered
}

// getRankingSortedEstablishments returns a sorted slice of establishments based on transaction amount for a given segment code
func getRankingSortedEstablishments(segmentCode int, segmentEstablishments map[int]map[int64]float64) []establishment {
	// Create a slice of establishments from the map and sort it by transaction amount in descending order
	var establishments []establishment

	// populate establishments slice with code and amount from segmentEstablishments for the given segment code
	for code, amount := range segmentEstablishments[segmentCode] {
		establishments = append(establishments, establishment{Code: code, Amount: amount})
	}

	// sort establishments by amount in descending order
	sort.Slice(establishments, func(i, j int) bool {
		return establishments[i].Amount > establishments[j].Amount
	})

	return establishments
}

// filterTopRanking filters the ranking data to include only the top 15 establishments based on transaction amount
func filterTopRanking(items map[string]*Ranking, segmentCode int, establishments []establishment) map[string]*RankingFiltered {
	// Group establishments by segment code and sum transaction amounts
	filtered := make(map[string]*RankingFiltered)

	// limits the count to topCount or the length of establishments if it's less than topCount
	count := topCount
	if len(establishments) < topCount {
		count = len(establishments)
	}

	// add greater establishments to filtered
	for i := 0; i < count; i++ {
		code := establishments[i].Code
		for key, ranking := range items {
			if ranking.SegmentCode == segmentCode && ranking.EstablishmentCode == code {
				filtered[key] = &RankingFiltered{
					Year:                ranking.Year,
					Quarter:             ranking.Quarter,
					EstablishmentCode:   ranking.EstablishmentCode,
					Function:            ranking.Function,
					Brand:               ranking.Brand,
					CaptureMode:         ranking.CaptureMode,
					Installments:        ranking.Installments,
					SegmentCode:         ranking.SegmentCode,
					TransactionAmount:   ranking.TransactionAmount,
					TransactionQuantity: ranking.TransactionQuantity,
					AvgMccFee:           ranking.AvgMccFee,
				}
			}
		}
	}

	// Return the filtered ranking data containing only the top establishments
	return filtered
}

// filterBottomRanking filters the ranking data to include only the bottom 200 establishments based on transaction amount
func filterBottomRanking(items map[string]*Ranking, segmentCode int, establishments []establishment) map[string]*RankingFiltered {
	// Group establishments by segment code and sum transaction amounts
	filtered := make(map[string]*RankingFiltered)

	// get last 200 establishments
	count := bottomCount
	if len(establishments) < count {
		count = len(establishments)
	}

	// filter bottom establishments and add to filtered_bottom
	filtered_bottom := make(map[string]*Ranking)
	for i := len(establishments) - count; i < len(establishments); i++ {
		code := establishments[i].Code
		for key, ranking := range items {
			if ranking.SegmentCode == segmentCode && ranking.EstablishmentCode == code {
				filtered_bottom[key] = &Ranking{
					Year:                ranking.Year,
					Quarter:             ranking.Quarter,
					EstablishmentCode:   ranking.EstablishmentCode,
					Function:            ranking.Function,
					Brand:               ranking.Brand,
					CaptureMode:         ranking.CaptureMode,
					Installments:        ranking.Installments,
					SegmentCode:         ranking.SegmentCode,
					TransactionAmount:   ranking.TransactionAmount,
					TransactionQuantity: ranking.TransactionQuantity,
					AvgMccFee:           ranking.AvgMccFee,
				}
			}
		}
	}

	// Return the filtered ranking data containing only the bottom establishments
	return filtered

}

// consolidateBottomRanking consolidates the bottom ranking data to include only the segment code, function, brand, capture mode, installments and segment code and sums the transaction amount and quantity and calculates the average mcc fee
func consolidateBottomRanking(filtered_bottom map[string]*RankingFiltered) map[string]*RankingFiltered {
	// Group establishments by segment code and sum transaction amounts
	filtered := make(map[string]*RankingFiltered)

	// consolidate last 200 establishments and add to filtered
	filtered_consolidate := consolidateRanking(filtered_bottom)
	for key, ranking := range filtered_consolidate {
		filtered[key] = &RankingFiltered{
			Year:                ranking.Year,
			Quarter:             ranking.Quarter,
			EstablishmentCode:   -1,
			Function:            ranking.Function,
			Brand:               ranking.Brand,
			CaptureMode:         ranking.CaptureMode,
			Installments:        ranking.Installments,
			SegmentCode:         ranking.SegmentCode,
			TransactionAmount:   ranking.TransactionAmount,
			TransactionQuantity: ranking.TransactionQuantity,
			AvgMccFee:           ranking.AvgMccFee,
		}
	}

	// Return the filtered ranking data containing only the bottom establishments
	return filtered
}

// Consolidate Ranking
func consolidateRanking(ranking map[string]*RankingFiltered) map[string]*RankingFiltered {
	// consolidate ranking by segment code, function, brand, capture mode, installments and segment code
	consRanking := make(map[string]*RankingFiltered)

	// for each ranking, if the establishment code is -1, consolidate it with the existing one in consRanking
	for _, r := range ranking {
		NewRanking := &RankingFiltered{
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

		// generate key for consRanking
		key := NewRanking.GetKey()

		// if the key already exists in consRanking, sum the transaction amount and quantity and calculate the new average mcc fee
		if existing, exists := consRanking[key]; exists {
			existing.TransactionAmount += r.TransactionAmount
			existing.TransactionQuantity += r.TransactionQuantity
			fee := existing.AvgMccFee / 100 * existing.TransactionAmount
			fee += r.AvgMccFee / 100 * r.TransactionAmount
			existing.AvgMccFee = fee / existing.TransactionAmount * 100
		} else {
			consRanking[key] = NewRanking
		}
	}

	return consRanking
}
