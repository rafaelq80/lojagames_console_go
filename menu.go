package main

import (
	"bufio"
	"fmt"
	"lojagames/model"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

var listaProdutos []model.Produto
var listaCategorias []model.Categoria
var codigoProduto = 0

func main() {

	reader := bufio.NewReader(os.Stdin)
	
	opcao := 0;

	var nome, fabricante string
	var preco float64
	var id, categoriaId int

	inicializaCategorias()

	for {

		fmt.Println()
		fmt.Println()
		fmt.Println("#######################################################")
		fmt.Println("                                                       ")
		fmt.Println("                LOJA DE GAMES GERAÇÃO Z                ")
		fmt.Println("                                                       ")
		fmt.Println("#######################################################")
		fmt.Println("                                                       ")
		fmt.Println("            1 - Criar Produto                          ")
		fmt.Println("            2 - Listar todos os Produtos               ")
		fmt.Println("            3 - Buscar Produto por Numero              ")
		fmt.Println("            4 - Buscar Produto por Nome                ")
		fmt.Println("            5 - Atualizar Produto                      ")
		fmt.Println("            6 - Apagar Produto                         ")
		fmt.Println("            7 - Sair                                   ")
		fmt.Println("                                                       ")
		fmt.Println("#######################################################")
		fmt.Println("Entre com a opção desejada:                            ")
		fmt.Println()

		fmt.Scan(&opcao)

		if opcao == 7 {
			sobre()
			os.Exit(0)
		}

		switch opcao {
		case 1:
			fmt.Println("Criar Produto")
			fmt.Println()

			fmt.Println("Digite o nome do Produto: ")
			nome,_ = reader.ReadString('\r')

			fmt.Println("Digite o nome do Fabricante do Produto: ")
			fabricante,_ = reader.ReadString('\r')

			fmt.Println("Digite o preço do Produto: ")
			fmt.Scanln(&preco)

			listarCategorias()
			fmt.Println("Informe o Id da Categoria do Produto: ")
			fmt.Scanln(&categoriaId)

			categoria := consultarCategoriaPorId(listaCategorias, categoriaId)

			if categoria != nil {
				produto := model.Produto{}
				produto.SetProduto(gerarId(), nome, fabricante, preco, *categoria)
				criarProduto(&listaProdutos, produto)
			}else{
				fmt.Println("Produto não encontrado!")
			}
		
			keyPress()
		case 2:
			fmt.Println("Listar todos os Produtos")
			fmt.Println()

			if len(listaProdutos) == 0{
				fmt.Println("Não existem produtos cadastrados")
			}else {
				listarProdutos(listaProdutos)
			}

			keyPress()
		case 3:
			fmt.Println("Consultar dados do produto - por Id")
			fmt.Println()

			fmt.Println("Digite o Id do Produto: ")
			fmt.Scan(&id)

			produto := consultarProdutosPorId(listaProdutos, id)

			if produto != nil {
				produto.Visualizar()
			}else{
				fmt.Println("Produto não encontrado!")
			}

			keyPress()
		case 4:
			fmt.Println("Consultar dados do produto - por Nome")
			fmt.Println()

			fmt.Println("Digite o nome do Produto: ")
			nome,_ = reader.ReadString('\r')

			filtroProdutos := consultarProdutosPorNome(listaProdutos, nome)

			for _, produto := range *filtroProdutos {
				produto.Visualizar()
			}

			keyPress()
		case 5:
			fmt.Println("Atualizar dados do produto")
			fmt.Println()

			fmt.Println("Digite o Id do Produto: ")
			fmt.Scan(&id)

			buscaProduto := consultarProdutosPorId(listaProdutos, id)

			if buscaProduto != nil {

				fmt.Println("Digite o nome do Produto: ")
				nome,_ = reader.ReadString('\r')

				fmt.Println("Digite o nome do Fabricante do Produto: ")
				fabricante,_ = reader.ReadString('\r')

				fmt.Println("Digite o preço do Produto: ")
				fmt.Scanln(&preco)

				fmt.Println("Informe a Categoria do Produto: ")
				fmt.Scanln(&categoriaId)

				categoria := consultarCategoriaPorId(listaCategorias, categoriaId)

				if categoria != nil {
					produtoAtualizado := model.Produto{}
					produtoAtualizado.SetProduto(id, nome, fabricante, preco, *categoria)
					produtoAtualizado.Visualizar()
					atualizarProduto(&listaProdutos, produtoAtualizado)
				}else{
					fmt.Println("Categoria não encontrada!")
				}

			}else{
				fmt.Println("Produto não encontrado!")
			}

			keyPress()
		case 6:
			fmt.Println("Apagar produto")
			fmt.Println()

			fmt.Println("Digite o Id do Produto: ")
			fmt.Scan(&id)

			buscaProduto := consultarProdutosPorId(listaProdutos, id)

			if buscaProduto != nil {
				deletarProduto(&listaProdutos, id)
			}else{
				fmt.Println("Produto não encontrado!")
			}

			keyPress()
		default:
			fmt.Println("Opção Inválida!")
			fmt.Println()

			keyPress()
		}
	}
}

func criarProduto(listaProdutos *[]model.Produto, produto model.Produto) {
	
	*listaProdutos = append(*listaProdutos, produto)
}

func listarProdutos(listaProdutos []model.Produto) {
	for _, produto := range listaProdutos {
		produto.Visualizar()
	}
}

func consultarProdutosPorId(listaProdutos []model.Produto, id int) (*model.Produto) {
	for _, produto := range listaProdutos {
		if(produto.GetId() == id){
			return &produto
		}
	}

	return nil
}

func consultarProdutosPorNome(listaProdutos []model.Produto, nome string) (*[]model.Produto) {
	var filtroProdutos []model.Produto
	
	for _, produto := range listaProdutos {
		if(strings.Contains(strings.TrimSpace(produto.GetNome()), strings.TrimSpace(nome))){
			filtroProdutos = append(filtroProdutos, produto)			
		}
	}

	return &filtroProdutos
}

func atualizarProduto(listaProdutos *[]model.Produto, produtoAtualizado model.Produto) {

	for i, p := range *listaProdutos {
		if p.GetId() == produtoAtualizado.GetId() {
			(*listaProdutos)[i] = produtoAtualizado
		}
	}

	fmt.Println()
	fmt.Println("O Produto id: ", produtoAtualizado.GetId(), " foi Atualizado!")
}

func deletarProduto(listaProdutos *[]model.Produto, id int) {
	for i, p := range *listaProdutos {
		if p.GetId() == id {
			*listaProdutos = append((*listaProdutos)[:i], (*listaProdutos)[i+1:]...)
		}
	}

	fmt.Println()
	fmt.Println("O Produto id: ", id, " foi Excluído!")
}

func gerarId() (int){
	codigoProduto ++
	return codigoProduto
}

func inicializaCategorias(){

	c1 := model.Categoria{}
	c1.SetCategoria(1, "Aventura")
	c2 := model.Categoria{}
	c2.SetCategoria(2, "E-Sports")
	
	listaCategorias = append(listaCategorias, c1)
	listaCategorias = append(listaCategorias, c2)

}

func listarCategorias(){

	fmt.Println()
	fmt.Println("Id - Descrição")
	for _, categoria := range listaCategorias {
		fmt.Println(categoria.GetId(), "-", categoria.GetDescricao())
	}
}

func consultarCategoriaPorId(listaCategorias []model.Categoria, categoriaId int) (*model.Categoria){
	for _, categoria := range listaCategorias {
		if(categoria.GetId() == categoriaId){
			return &categoria
		}
	}

	return nil
}

func keyPress()(rune){
	fmt.Println()
	fmt.Println("Pressione enter para continuar")
	char, _, err := keyboard.GetSingleKey()
	if (err != nil) {
		panic(err)
	}

	return char
}

func sobre(){
	fmt.Println()
	fmt.Println("##################################################")
	fmt.Println("Lojas Geração Z                                   ")
	fmt.Println("##################################################")
	fmt.Println("Rafael Queiróz                                    ")
	fmt.Println("rafaelproinfo@gmail.com                           ")
	fmt.Println("github.com/rafaelq80                              ")
	fmt.Println("##################################################")
	fmt.Println()
}