package validator

import (
	"fmt"
	"math"

	"conversorCadoc9011/internal/excel"
)

func ValidateHierarchy(rows []*excel.RawRow, periods []string) error {

	for _, parent := range rows {

		if parent.Nivel == "" {
			continue
		}

		for _, period := range periods {

			sum := 0.0
			hasChildren := false

			for _, child := range rows {

				if child.Demonstrativo != parent.Demonstrativo {
					continue
				}

				if child.PaiNivel != parent.Nivel {
					continue
				}

				sum += child.Valores[period]
				hasChildren = true
			}

			if !hasChildren {
				continue
			}

			parentValue := parent.Valores[period]
			diff := sum - parentValue

			if math.Abs(diff) > 0.01 {

				return fmt.Errorf(`
					ERRO NA PLANILHA

					Conta pai: %s
					Linha Excel: %d
					Periodo: %s

					Valor informado no pai: %.2f
					Soma dos filhos: %.2f
					Diferença encontrada: %.2f

					Verifique os valores das contas filhas dessa conta.
					`,
					parent.Conta,
					parent.ExcelRow,
					period,
					parentValue,
					sum,
					diff,
				)
			}
		}
	}

	return nil
}
