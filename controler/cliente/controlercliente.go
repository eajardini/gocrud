package controlercliente

import (
	bd "eajardini/gin/gocrud/bancodados"
	modelCliente "eajardini/gin/gocrud/controler/cliente/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	bancodados bd.BDCon
)

//OlaCliente : zz
func OlaCliente(c *gin.Context) {

	mensagem := gin.H{
		"resposta": "Olá Cliente!",
	}

	c.JSON(200, mensagem)
}

//InsereCliente :zz
func InsereCliente(c *gin.Context) {
	//Para testar:
	//curl --header "Content-Type: application/json" --request POST  --data '{"clienteID":121,"nome":"Borodin", "endereco":"Rua 15","datanasc":"21/11/2019"}' http://localhost:8080/cliente/insere

	var clienteJSON modelCliente.TipoCliente

	bancodados.AbreConexao()
	defer bancodados.FechaConexao()

	erro := c.BindJSON(&clienteJSON)

	if erro != nil {
		log.Fatal("[controlecliente, InsereCliente ] Erro ao fazer Bind de Json!!")
	}

	sql := `INSERT INTO cliente (clienteID, nome, endereco, datanasc ) values ($1,$2,$3,$4)`

	layoutDaData := "02/01/2006" //DD/MM/YYYY
	datanasc, _ := time.Parse(layoutDaData, clienteJSON.Datanasc)
	tx := bancodados.BD.MustBegin()
	_, err := tx.Exec(sql, clienteJSON.ClienteID, clienteJSON.Nome, clienteJSON.Endereco, datanasc)

	if err != nil {
		resposta := gin.H{
			"resposta": err.Error(),
		}
		tx.Rollback()
		c.JSON(200, resposta)
		return
	}
	tx.Commit()

	mensagem := gin.H{
		"resposta": "Cliente inserido com sucesso!",
	}

	c.JSON(200, mensagem)
}

//SelecionaTodosOsCliente :zz
func SelecionaTodosOsCliente(c *gin.Context) {
	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  http://localhost:8080/cliente/selecionatodos

	var (
		clienteBD []modelCliente.TipoCliente
		// clienteJSON modelCliente.TipoCliente
	)

	bancodados.AbreConexao()
	defer bancodados.FechaConexao()

	sql := `SELECT * FROM cliente ORDER BY nome`
	err := bancodados.BD.Select(&clienteBD, sql)

	if err != nil {
		resposta := gin.H{
			"resposta": err.Error(),
		}
		c.JSON(200, resposta)
		return
	}

	c.JSON(200, clienteBD)
}

//SelecionaClientePorNome :zz
func SelecionaClientePorNome(c *gin.Context) {
	//Para testar:
	//curl --header "Content-Type: application/json" --request GET  --data '{"nomeParaBusca":"Abel"}' http://localhost:8080/cliente/selecionaclientepornome

	type tipoNomeParaBusca struct {
		NomeParaBusca string `json:"nomeParaBusca"`
	}

	var (
		clienteBD     []modelCliente.TipoCliente
		nomeParaBusca tipoNomeParaBusca
	)

	bancodados.AbreConexao()
	defer bancodados.FechaConexao()

	erro := c.BindJSON(&nomeParaBusca)

	if erro != nil {
		log.Fatal("[controlecliente, SelecionaClientePorNome ] Erro ao fazer Bind de Json!!")
	}

	sql := `SELECT * 
					FROM cliente 
					where nome like $1
					ORDER BY nome`

	err := bancodados.BD.Select(&clienteBD, sql, nomeParaBusca.NomeParaBusca)

	if err != nil {
		resposta := gin.H{
			"resposta": err.Error(),
		}
		c.JSON(200, resposta)
		return
	}

	c.JSON(200, clienteBD)
}

//AtualizaCliente :zz
func AtualizaCliente(c *gin.Context) {
	//Para testar:
	//curl --header "Content-Type: application/json" --request PUT  --data '{"clienteID":121,"nome":"Borodin", "endereco":"Rua 15","datanasc":"21/11/2019"}' http://localhost:8080/cliente/atualizacliente

	var (
		clienteJSON modelCliente.TipoCliente
	)

	bancodados.AbreConexao()
	defer bancodados.FechaConexao()

	erro := c.BindJSON(&clienteJSON)

	if erro != nil {
		log.Fatal("[controlecliente, AtualizaClientePorNome ] Erro ao fazer Bind de Json!!")
	}

	sql := `update cliente 
					set nome = $2,
							endereco = $3,
							datanasc = $4
					where clienteID = $1`

	layoutDaData := "02/01/2006" //DD/MM/YYYY
	datanasc, _ := time.Parse(layoutDaData, clienteJSON.Datanasc)
	tx := bancodados.BD.MustBegin()
	sqlResult, err := tx.Exec(sql, clienteJSON.ClienteID, clienteJSON.Nome, clienteJSON.Endereco, datanasc)
	if err != nil {
		resposta := gin.H{
			"resposta": err.Error(),
		}
		c.JSON(200, resposta)
		return
	}
	linhasAtualizadas, _ := sqlResult.RowsAffected()
	if linhasAtualizadas == 0 {
		resposta := gin.H{
			"resposta": "Não foi atualizado nenhum cliente",
		}
		tx.Rollback()
		c.JSON(200, resposta)
		return
	}

	tx.Commit()

	resposta := gin.H{
		"resposta": "Cliente atualizado com sucesso",
	}

	c.JSON(200, resposta)
}

//ApagaCliente :zz
func ApagaCliente(c *gin.Context) {
	//Para testar:
	//curl --header "Content-Type: application/json" --request DELETE  --data '{"clienteID":121}' http://localhost:8080/cliente/apagacliente
	type tipoClienteIDParaApagar struct {
		ClienteID int `json:"clienteID"`
	}

	var (
		clienteIDParaApagar tipoClienteIDParaApagar
	)

	bancodados.AbreConexao()
	defer bancodados.FechaConexao()

	erro := c.BindJSON(&clienteIDParaApagar)

	if erro != nil {
		log.Fatal("[controlecliente, ApagaCliente ] Erro ao fazer Bind de Json!!")
	}

	sql := `delete from cliente
					where clienteID = $1`

	tx := bancodados.BD.MustBegin()
	sqlResult, err := tx.Exec(sql, clienteIDParaApagar.ClienteID)

	if err != nil {
		resposta := gin.H{
			"resposta": err.Error(),
		}
		c.JSON(200, resposta)
		tx.Rollback()
		return
	}

	linhasApagadas, _ := sqlResult.RowsAffected()

	if linhasApagadas == 0 {
		resposta := gin.H{
			"resposta": "Nenhum cliente foi apagado",
		}

		c.JSON(200, resposta)
		tx.Rollback()
		return
	}

	tx.Commit()

	mensagem := gin.H{
		"resposta": "Cliente apagado com sucesso",
	}

	c.JSON(200, mensagem)
}
