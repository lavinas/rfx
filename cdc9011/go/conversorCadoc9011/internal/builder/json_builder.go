package builder

import (
	"encoding/json"
	"fmt"
	"os"

	"conversorCadoc9011/internal/excel"
	"conversorCadoc9011/internal/model"
)

func BuildDocument(meta excel.Meta, rows []*excel.RawRow, periods []string) (model.Documento, error) {

	doc := model.Documento{
		CNPJ:            meta.CNPJ,
		CodigoDocumento: meta.CodigoDocumento,
		TipoRemessa:     meta.TipoRemessa,
		UnidadeMedida:   meta.UnidadeMedida,
		DataBase:        meta.DataBase,
	}

	for i, p := range periods {

		doc.DatasBaseReferencia = append(
			doc.DatasBaseReferencia,
			model.DtRef{
				ID:   fmt.Sprintf("dt%d", i+1),
				Data: p,
			},
		)
	}

	dateToID := make(map[string]string)
	for i, p := range periods {
		dateToID[p] = fmt.Sprintf("dt%d", i+1)
	}

	for _, r := range rows {

		c := model.Conta{
			ID:        r.ID,
			Nivel:     r.Nivel,
			Descricao: r.Conta,
			ContaPai:  r.ContaPaiID,
		}

		for _, p := range periods {

			v, ok := r.Valores[p]
			if !ok {
				return doc, fmt.Errorf(
					"valor não encontrado para conta %s periodo %s",
					r.Conta,
					p,
				)
			}

			c.ValoresIndividualizados = append(
				c.ValoresIndividualizados,
				model.Valor{
					DtBase: dateToID[p],
					Valor:  v,
				},
			)
		}

		switch r.Demonstrativo {

		case "BalancoPatrimonial":

			doc.BalancoPatrimonial.Contas =
				append(doc.BalancoPatrimonial.Contas, c)

		case "DemonstracaoDoResultado":

			doc.DemonstracaoDoResultado.Contas =
				append(doc.DemonstracaoDoResultado.Contas, c)

		default:

			return doc, fmt.Errorf(
				"demonstrativo desconhecido: %s (linha %d)",
				r.Demonstrativo,
				r.ExcelRow,
			)
		}
	}

	return doc, nil
}

func WriteJSON(path string, doc model.Documento) error {

	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
