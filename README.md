<h1 align="center">
Go: Crie uma aplicação web
</h1>

### Índice

- Servidor e struct de produtos 
- Conectando com banco de dados
- Refatoração e página de novos produtos
- Deletando produtos e partials
- Atualizando e editando produtos

<h2 align="center">
Servidor e struct de produtos 
</h2>

* Criamos o nosso projeto no workspace correto, dentro do GOPATH (dentro da pasta src, github.com, seguido do nome de usuário do Github);

* Aprendemos como subir um servidor através da função http.ListenAndServe(), exibindo uma tabela com nossos produtos;

* Em seguida criamos uma struct de Produto, onde instanciamos alguns produtos e exibimos de forma dinâmica em nossa index.html.

<h2 align="center">
Conectando com banco de dados
</h2>

* Instalamos o Postgres para armazenar nossos produtos de forma segura;

* Criamos uma função chamada conectaComBancoDeDados() para abrir a conexão com o banco de dados;

* Alteramos nosso código para exibir os produtos que estão cadastrados lá no banco de dados.

Script utilizado para criação da tabela.:

```sql
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)
```
Script utilizado para inserção de itens na tabela produtos.:
```sql 
insert into produtos (nome, descricao, preco, quantidade) values 
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito Bom', 99, 5);

```

<h2 align="center">
Refatoração e página de novos produtos
</h2>

* Modularizamos o código para tornar a manutenção e/ou atualização mais clara, criando as pastas controllers, db, models, routes e templates;

* Criamos uma página para criar novos produtos e uma rota capaz de atender essa requisição http.HandleFunc("/new", controllers.New);

* Buscamos os dados da página new com o código r.FormValue() para cada input (nome, descrição, preço e quantidade) no controller de produtos;

* Salvamos o produto através do modelo de produto criando a função CriaNovoProduto().

<h2 align="center">
Deletando produtos e partials
</h2>

* Criamos um botão na linha de cada produto que assim que clicado, deletava o produto do banco de dados;

* Para melhorar a remoção dos produtos, criamos uma função em Javascript perguntando se queremos de fato, deletar o produto;

* Removemos o código HTML duplicado da index e do arquivo new, criando as seguintes partials: _head.html e _menu.html.

<h2 align="center">
Atualizando e editando produtos
</h2>

* Criamos um botão na linha de cada produto que nos direciona para a tela de edição;

* Após criar a tela de edição, preenchemos o formulário com as informações do produto exibindo os dados já cadastrados;

* Atualizamos o produto, buscando os dados alterados na tela e executando o update(atualização) no banco de dados.

### 🛠 Tecnologias

- [GoLang 1.20](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/download/)


### 🛠  Ferramentas:

- [Compilador Online](https://go.dev/play/p/gkwKo7rholt)

