package excel

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Meta struct {
	CNPJ            string
	CodigoDocumento string
	TipoRemessa     string
	UnidadeMedida   int
	DataBase        string
}

type RawRow struct {
	ExcelRow int

	Demonstrativo string
	Nivel         string
	Conta         string
	PaiNivel      string

	Valores map[string]float64

	ID         string
	ContaPaiID string
}

func ParseFile(path string) (Meta, []*RawRow, []string, error) {

	file, err := excelize.OpenFile(path)
	if err != nil {
		return Meta{}, nil, nil, err
	}

	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		return Meta{}, nil, nil, err
	}

	meta, headerIndex, err := readMetadata(rows)
	if err != nil {
		return Meta{}, nil, nil, err
	}

	periods, periodStart := detectPeriods(rows[headerIndex])
	dataRows := parseRows(rows, headerIndex+1, periods, periodStart)

	assignIDs(dataRows)

	resolveParents(dataRows)

	return meta, dataRows, periods, nil
}

func readMetadata(rows [][]string) (Meta, int, error) {

	meta := Meta{}
	header := -1

	for i, row := range rows {

		if len(row) == 0 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(row[0]))

		switch key {

		case "cnpj":
			if len(row) > 1 {
				meta.CNPJ = row[1]
			}

		case "codigodocumento":
			meta.CodigoDocumento = row[1]

		case "tiporemessa":
			meta.TipoRemessa = row[1]

		case "unidademedida":
			meta.UnidadeMedida, _ = strconv.Atoi(row[1])

		case "database":
			meta.DataBase = row[1]

		case "demonstrativo":
			header = i
		}
	}

	if header == -1 {
		return meta, -1, fmt.Errorf("header não encontrado")
	}

	return meta, header, nil
}

func detectPeriods(header []string) ([]string, int) {

	var periods []string
	start := -1

	for i, h := range header {

		h = strings.TrimSpace(h)

		if strings.HasPrefix(h, "A") {

			if start == -1 {
				start = i
			}

			periods = append(periods, h)
		}
	}

	return periods, start
}
