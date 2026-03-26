package model

type DtRef struct {
	ID   string `json:"@id"`
	Data string `json:"@data"`
}

type Valor struct {
	DtBase string  `json:"@dtBase"`
	Valor  float64 `json:"-"`
	ValorCurrence Currence `json:"@valor"`
}

type Conta struct {
	ID        string `json:"@id"`
	Nivel     string `json:"@nivel"`
	Descricao string `json:"@descricao"`
	ContaPai  string `json:"@contaPai"`

	ValoresIndividualizados []Valor `json:"valoresIndividualizados"`
}

type Bloco struct {
	Contas []Conta `json:"contas,omitempty"`
}

type Documento struct {
	CNPJ            string `json:"@cnpj"`
	CodigoDocumento string `json:"@codigoDocumento"`
	TipoRemessa     string `json:"@tipoRemessa"`
	UnidadeMedida   int    `json:"@unidadeMedida"`
	DataBase        string `json:"@dataBase"`

	DatasBaseReferencia []DtRef `json:"datasBaseReferencia"`

	BalancoPatrimonial                                            *Bloco `json:"BalancoPatrimonial,omitempty"`
	DemonstracaoDoResultado                                       *Bloco `json:"DemonstracaoDoResultado,omitempty"`
	DemonstracaoDoResultadoAbrangente                             *Bloco `json:"DemonstracaoDoResultadoAbrangente,omitempty"`
	DemonstracaoDosFluxosDeCaixa                                  *Bloco `json:"DemonstracaoDosFluxosDeCaixa,omitempty"`
	DemonstracaoDasMutacoesDoPatrimonioLiquido                    *Bloco `json:"DemonstracaoDasMutacoesDoPatrimonioLiquido,omitempty"`
	DemonstracaoDosRecursosDeConsorcioConsolidada                 *Bloco `json:"DemonstracaoDosRecursosDeConsorcioConsolidada,omitempty"`
	DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada *Bloco `json:"DemonstracaoDeVariacoesNasDisponibilidadesDeGruposConsolidada,omitempty"`
}
