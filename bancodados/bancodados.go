package bancodados

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//BDCon : zz
type BDCon struct {
	BD *sqlx.DB
}

//IniciaConexao : zz
func (bdcon *BDCon) IniciaConexao() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=crudbd host=172.17.0.1  sslmode=disable")
	if err != nil {
		log.Fatal("[bancodados, Inicia] Erro ao Conectar ao Banco de Dados!!")
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("[bancodados, Inicia] Erro ao Pingar ao Banco de Dados!!")
	} else {
		log.Println("[bancodados, Inicia] Banco de Dados conectado corretamente!!")
	}

	db.Close()
}

//AbreConexao :zz
func (bdcon *BDCon) AbreConexao() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=crudbd host=172.17.0.1  sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	//seta a variavel BD para fazer as consultas SQL
	bdcon.BD = db

}

//FechaConexao :zz
func (bdcon *BDCon) FechaConexao() {
	bdcon.BD.Close()
}

//Insert : faz o insert no Banco de Dados
func (bdcon *BDCon) Insert(sgbd string, sqlinsert interface{}) {
	fmt.Println(sqlinsert)
}
