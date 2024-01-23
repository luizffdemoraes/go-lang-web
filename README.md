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

### 🛠 Tecnologias

- [GoLang 1.20](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/download/)


### 🛠  Ferramentas:

- [Compilador Online](https://go.dev/play/p/gkwKo7rholt)

