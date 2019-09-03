package main

import(
	"log"
	"fmt"
	"github.com/lucashtc/desafio-dito/coletora"
	"github.com/pkg/errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	con,err := coletora.Conn()
	if err != nil {
		log.Fatal(errors.Wrap(err,"falha ao conectar ao banco"))
	}

	col := coletora.Colector{}
	col.Event = "Teste"
	col.Timestamp = "Teste"

	err = coletora.Insert(con,col)
	if err != nil {
		log.Fatal("Deu ruim")
	}
	fmt.Println("Deu")


}