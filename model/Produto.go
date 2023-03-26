package model

import "fmt"

type Produto struct {
	id         int
	nome       string
	fabricante string
	preco      float64
	categoria  Categoria
}

func (p *Produto) SetProduto(id int, nome string, fabricante string, preco float64, categoria Categoria) {
	p.SetId(id)
	p.SetNome(nome)
	p.SetFabricante(fabricante)
	p.SetPreco(preco)
	p.categoria.SetCategoria(categoria.GetId(), categoria.GetDescricao())
}

func (p *Produto) SetId(id int) {
	p.id = id
}

func (p *Produto) GetId() int {
	return p.id
}

func (p *Produto) SetNome(nome string) {
	p.nome = nome
}

func (p *Produto) GetNome() string {
	return p.nome
}

func (p *Produto) SetFabricante(fabricante string) {
	p.fabricante = fabricante
}

func (p *Produto) GetFabricante() string {
	return p.fabricante
}

func (p *Produto) SetPreco(preco float64) {
	p.preco = preco
}

func (p *Produto) GetPreco() float64 {
	return p.preco
}

func (p *Produto) Visualizar() {

	fmt.Println()
	fmt.Println()
	fmt.Println("******************************************")
	fmt.Println("Dados do Produto")
	fmt.Println("******************************************")
	fmt.Println("Id: ", p.GetId())
	fmt.Println("Nome: ", p.GetNome())
	fmt.Println("Fabricante: ", p.GetFabricante())
	fmt.Println("Preço: ", p.GetPreco())
	fmt.Println("Descrição: ", p.categoria.GetDescricao())
}
