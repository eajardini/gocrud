package modelcliente

//TipoCliente :zz
type TipoCliente struct {
	ClienteID int    `json:"clienteID"`
	Nome      string `json:"nome"`
	Endereco  string `json:"endereco"`
	Datanasc  string `json:"datanasc"`
}
