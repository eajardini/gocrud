package route

import (
	bd "eajardini/gin/gocrud/bancodados"
	cliente "eajardini/gin/gocrud/controler/cliente"

	"github.com/gin-gonic/gin"
)

func estaEmFuncionamento(c *gin.Context) {

	c.JSON(200, gin.H{
		"resposta": "Está funcionando!",
	})

}

var bancodados bd.BDCon

//IniciaServidor : sd
func IniciaServidor() {
	r := gin.Default()

	bancodados.IniciaConexao()

	r.GET("/", estaEmFuncionamento)

	cli := r.Group("/cliente")
	{
		cli.GET("", cliente.OlaCliente)
		cli.GET("/selecionatodos", cliente.SelecionaTodosOsCliente)
		cli.GET("/selecionaclientepornome", cliente.SelecionaClientePorNome)
		cli.POST("/insere", cliente.InsereCliente)
		cli.PUT("/atualizacliente", cliente.AtualizaCliente)
		cli.DELETE("/apagacliente", cliente.ApagaCliente)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
