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

func CriarNovoProduto(nome, descricao string, quantidade int, valor float64) {
	db := conexao.Conexao()
	prepare, err := db.Prepare("insert into produtos (nome, descricao, quantidade, valor) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(nome, descricao, quantidade, valor)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := conexao.Conexao()
	query, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoAtualizar := Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err := query.Scan(&id, &nome, &descricao, &quantidade, &valor)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Quantidade = quantidade
		produtoAtualizar.Valor = valor
	}
	defer db.Close()
	return produtoAtualizar
}

func AtualizaProduto(id int, nome, descricao string, quantidade int, valor float64) {
	db := conexao.Conexao()
	prepare, err := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, valor=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(nome, descricao, quantidade, valor, id)
	defer db.Close()
}

func DeleteProduto(idProduto string) {
	db := conexao.Conexao()
	prepare, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec(idProduto)
	defer db.Close()
}
