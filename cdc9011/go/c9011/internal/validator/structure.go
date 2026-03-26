package validator

import (
	"fmt"
	"c9011/internal/model"

)


func ValidateStructure(doc model.Documento) error {
	err := ""

	if doc.BalancoPatrimonial == nil {
		err += "BalancoPatrimonial está ausente. "
	}
	if doc.DemonstracaoDoResultado == nil {
		err += "DemonstracaoDoResultado está ausente. "
	}
	if doc.DemonstracaoDoResultadoAbrangente == nil {
		err += "DemonstracaoDoResultadoAbrangente está ausente. "
	}
	if doc.DemonstracaoDosFluxosDeCaixa == nil {
		err += "DemonstracaoDosFluxosDeCaixa está ausente. "
	}
	if doc.DemonstracaoDasMutacoesDoPatrimonioLiquido == nil {
		err += "DemonstracaoDasMutacoesDoPatrimonioLiquido está ausente. "
	}
	if err != "" {
		return fmt.Errorf("ERRO NA PLANILHA: %s", err)
	}
	return nil
}