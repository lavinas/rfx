package target_domain

import (
	"sort"

	"consolidator/internal/core/ports"
)

const (
	// TopCount defines the number of top establishments to include in the filtered ranking
	topCount = 15
	// BottomCount defines the number of bottom establishments to include in the filtered ranking
	bottomCount = 200
)

// RankingsFiltered represents the filtered ranking of establishments based on transaction data.
type RankingFiltered struct {
	rankingItem *RankingFilteredItem
	consolidation map[string]*RankingFilteredItem
}

// establishment represents the data structure for establishments which will be used for filtering the ranking data
type establishment struct {
	Code   int64
	Amount float64
}

// NewRankingFiltered creates a new instance of RankingFiltered.
func NewRankingFiltered() *RankingFiltered {
	return &RankingFiltered{
		rankingItem: NewRankingFilteredItem(),
		consolidation: make(map[string]*RankingFilteredItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (r *RankingFiltered) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&RankingFilteredItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (r *RankingFiltered) Save(repository ports.Repository) error {
	if err := repository.Save(r.consolidation); err != nil {
		return err
	}
	return nil
}


// FilterRanking filters the ranking data to include only the top and bottom establishments based on transaction amount
func (r *RankingFiltered) FilterRanking(items map[string]*RankingItem) {

	// Group establishments by segment code and sum transaction amounts
	segmentEstablishments := r.getSegmentEstablishmentsMap(items)

	// for each segment code, filter top and bottom establishments
	for segmentCode := range segmentEstablishments {
		// sort establishments by transaction amount
		establishments := r.getRankingSortedEstablishments(segmentCode, segmentEstablishments)

		// add top establishments
		r.addTopRanking(items, segmentCode, establishments)

		// add bottom establishments
		r.addBottomRanking(items, segmentCode, establishments)

	}
}

// getSegmentEstablishments returns a map of segment code to a map of establishment code and transaction amount
func (r *RankingFiltered) getSegmentEstablishmentsMap(items map[string]*RankingItem) map[int]map[int64]float64 {
	segmentEstablishments := make(map[int]map[int64]float64)
	for _, ranking := range items {
		if _, exists := segmentEstablishments[ranking.SegmentCode]; !exists {
			segmentEstablishments[ranking.SegmentCode] = make(map[int64]float64)
		}
		segmentEstablishments[ranking.SegmentCode][ranking.EstablishmentCode] += ranking.TransactionAmount
	}
	return segmentEstablishments
}

// getRankingSortedEstablishments returns a sorted slice of establishments based on transaction amount for a given segment code
func (r *RankingFiltered) getRankingSortedEstablishments(segmentCode int, segmentEstablishments map[int]map[int64]float64) []establishment {
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

// addTopRanking adds the top 15 establishments based on transaction amount to the consolidation
func (r *RankingFiltered) addTopRanking(items map[string]*RankingItem, segmentCode int, establishments []establishment) {
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
				r.consolidation[key] = &RankingFilteredItem{
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
}


// addBottomRanking2 adds the bottom  establishments based on transaction amount to the consolidation
func (r *RankingFiltered) addBottomRanking(items map[string]*RankingItem, segmentCode int, establishments []establishment) {
	// get bottom establishments
	count := bottomCount
	if len(establishments) < count {
		count = len(establishments)
	}
	// filter bottom establishments and add to filtered_bottom
	filteredRanking := map[string]*RankingItem{}
	for i := len(establishments) - count; i < len(establishments); i++ {
		code := establishments[i].Code
		for key, ranking := range items {
			if ranking.SegmentCode == segmentCode && ranking.EstablishmentCode == code {
				filteredRanking[key] = ranking
			}
		}
	}
	// consolidate the filtered bottom ranking data	
	r.consolidateBottomRanking(filteredRanking)
}

// Consolidate Ranking
func (r *RankingFiltered) consolidateBottomRanking(ranking map[string]*RankingItem) {

	// for each ranking, if the establishment code is -1, consolidate it with the existing one in consRanking
	for _, rk := range ranking {
		NewRanking := &RankingFilteredItem{
			Year:                rk.Year,
			Quarter:             rk.Quarter,
			EstablishmentCode:   -1,
			Function:            rk.Function,
			Brand:               rk.Brand,
			CaptureMode:         rk.CaptureMode,
			Installments:        rk.Installments,
			SegmentCode:         rk.SegmentCode,
			TransactionAmount:   rk.TransactionAmount,
			TransactionQuantity: rk.TransactionQuantity,
			AvgMccFee:           rk.AvgMccFee,
		}

		// generate key for consRanking
		key := NewRanking.GetKey()

		// if the key already exists in consRanking, sum the transaction amount and quantity and calculate the new average mcc fee
		if existing, exists := r.consolidation[key]; exists {
			fee := existing.AvgMccFee / 100 * existing.TransactionAmount
			fee += rk.AvgMccFee / 100 * rk.TransactionAmount
			existing.AvgMccFee = fee / (existing.TransactionAmount + rk.TransactionAmount) * 100
			existing.TransactionAmount += rk.TransactionAmount
			existing.TransactionQuantity += rk.TransactionQuantity
		} else {
			r.consolidation[key] = NewRanking
		}
	}
}
