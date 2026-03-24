package excel

import (
	"fmt"

	"c9011/internal/util"
)

func parseRows(rows [][]string, start int, periods []string, periodStart int) []*RawRow {

	var result []*RawRow
	currentDemo := ""

	for i := start; i < len(rows); i++ {

		row := rows[i]

		demo := util.GetCell(row, 0)
		if demo != "" {
			currentDemo = demo
		}

		nivel := util.GetCell(row, 1)
		conta := util.GetCell(row, 3)

		if nivel == "" && conta == "" {
			continue
		}

		r := &RawRow{
			ExcelRow:      i + 1,
			Demonstrativo: currentDemo,
			Nivel:         nivel,
			Conta:         conta,
			PaiNivel:      util.NormalizeParent(util.GetCell(row, 2)),
			Valores:       map[string]float64{},
		}

		for pIndex, p := range periods {

			col := periodStart + pIndex
			valor := util.ParseFloat(util.GetCell(row, col))

			r.Valores[p] = valor
		}

		result = append(result, r)
	}

	return result
}

func assignIDs(rows []*RawRow) {

	for i, r := range rows {

		r.ID = fmt.Sprintf("conta%d", i+1)
	}
}

func resolveParents(rows []*RawRow) {

	lastSeenLevel := map[string]string{}

	for _, r := range rows {

		if r.PaiNivel != "" && r.PaiNivel != "-" {

			parentKey := r.Demonstrativo + "|" + r.PaiNivel

			if parentID, ok := lastSeenLevel[parentKey]; ok {
				r.ContaPaiID = parentID
			}
		}

		key := r.Demonstrativo + "|" + r.Nivel
		lastSeenLevel[key] = r.ID
	}
}
