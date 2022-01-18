package models

import conexao "go-web-application/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Quantidade      int
	Valor           float64
}

func BuscaTodosProdutos() []Produto {

	db := conexao.Conexao()
	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &quantidade, &valor)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Valor = valor

		produtos = append(produtos, p)
		defer db.Close()
	}

	return produtos
}
