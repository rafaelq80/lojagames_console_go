package model

import "fmt"

type Categoria struct {
	id         int
	descricao  string
}

func (p *Categoria) SetCategoria(id int, descricao string) {
	p.SetId(id)
	p.SetDescricao(descricao)
}

func (p *Categoria) SetId(id int) {
	p.id = id
}

func (p *Categoria) GetId() int {
	return p.id
}

func (p *Categoria) SetDescricao(descricao string) {
	p.descricao = descricao
}

func (p *Categoria) GetDescricao() string {
	return p.descricao
}

func (p *Categoria) Visualizar() {

	fmt.Println()
	fmt.Println()
	fmt.Println("******************************************")
	fmt.Println("Dados do Categoria")
	fmt.Println("******************************************")
	fmt.Println("Id do Categoria: ", p.GetId())
	fmt.Println("descricao do Categoria: ", p.GetDescricao())

}
