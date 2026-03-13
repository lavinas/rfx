package model

type DtRef struct {
	ID   string `json:"@id"`
	Data string `json:"@data"`
}

type Valor struct {
	DtBase string  `json:"@dtBase"`
	Valor  float64 `json:"@valor"`
}

type Conta struct {
	ID        string `json:"@id"`
	Nivel     string `json:"@nivel"`
	Descricao string `json:"@descricao"`
	ContaPai  string `json:"@contaPai,omitempty"`

	ValoresIndividualizados []Valor `json:"valoresIndividualizados,omitempty"`
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

	BalancoPatrimonial      Bloco `json:"balancoPatrimonial"`
	DemonstracaoDoResultado Bloco `json:"demonstracaoDoResultado"`
}
