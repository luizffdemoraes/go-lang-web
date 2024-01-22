package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB{
	conexao := "user=postgres dbname=alura_loja password= host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("loja/templates/*.html"))

func main() {
	db := conectaComBancoDeDados();
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{Nome: "Tenis", Descricao: "Confortável", Preco: 89, Quantidade: 5},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59, Quantidade: 2},
		{Nome: "Produto Novo", Descricao: "Muito legal", Preco: 1.99, Quantidade: 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
