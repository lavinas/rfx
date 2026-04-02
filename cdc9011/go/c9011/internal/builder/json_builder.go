package builder

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"c9011/internal/excel"
	"c9011/internal/model"
	"c9011/internal/util"
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

			if math.IsNaN(v) {
				continue
			}

			c.ValoresIndividualizados = append(
				c.ValoresIndividualizados,
				model.Valor{
					DtBase:        dateToID[p],
					Valor:         util.Round2(v),
					ValorCurrence: model.Currence(util.Round2(v)),
				},
			)
		}

		switch r.Demonstrativo {

		case "BalancoPatrimonial":
			if doc.BalancoPatrimonial == nil {
				doc.BalancoPatrimonial = &model.Bloco{}
			}

			doc.BalancoPatrimonial.Contas =
				append(doc.BalancoPatrimonial.Contas, c)

		case "DemonstracaoDoResultado":
			if doc.DemonstracaoDoResultado == nil {
				doc.DemonstracaoDoResultado = &model.Bloco{}
			}

			doc.DemonstracaoDoResultado.Contas =
				append(doc.DemonstracaoDoResultado.Contas, c)

		case "DemonstracaoDoResultadoAbrangente":
			if doc.DemonstracaoDoResultadoAbrangente == nil {
				doc.DemonstracaoDoResultadoAbrangente = &model.Bloco{}
			}

			doc.DemonstracaoDoResultadoAbrangente.Contas =
				append(doc.DemonstracaoDoResultadoAbrangente.Contas, c)

		case "DemonstracaoDosFluxosDeCaixa":
			if doc.DemonstracaoDosFluxosDeCaixa == nil {
				doc.DemonstracaoDosFluxosDeCaixa = &model.Bloco{}
			}

			doc.DemonstracaoDosFluxosDeCaixa.Contas =
				append(doc.DemonstracaoDosFluxosDeCaixa.Contas, c)

		case "DemonstracaoDasMutacoesDoPatrimonioLiquido":
			if doc.DemonstracaoDasMutacoesDoPatrimonioLiquido == nil {
				doc.DemonstracaoDasMutacoesDoPatrimonioLiquido = &model.Bloco{}
			}

			doc.DemonstracaoDasMutacoesDoPatrimonioLiquido.Contas =
				append(doc.DemonstracaoDasMutacoesDoPatrimonioLiquido.Contas, c)

		case "DemonstracaoDosRecursosDeConsorcioConsolidada":
			if doc.DemonstracaoDosRecursosDeConsorcioConsolidada == nil {
				doc.DemonstracaoDosRecursosDeConsorcioConsolidada = &model.Bloco{}
			}

			doc.DemonstracaoDosRecursosDeConsorcioConsolidada.Contas =
				append(doc.DemonstracaoDosRecursosDeConsorcioConsolidada.Contas, c)

		case "DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada":
			if doc.DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada == nil {
				doc.DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada = &model.Bloco{}
			}

			doc.DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada.Contas =
				append(doc.DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada.Contas, c)

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
