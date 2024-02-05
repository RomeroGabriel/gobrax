# Gobrax Challenge

Este repositório visa solucionar o desafio de Backend challenge, que inclui as seguintes funcionalidades:

1. CRUD de motoristas (Criação, Listagem, Atualização e Remoção)
1. CRUD de veículos (Criação, Listagem, Atualização e Remoção)
1. Vinculação de motorista a um veículo

## Estrutura

Para estruturar e organizar a aplicação, foram seguidas algumas recomendações do [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Ferramentas

Para executar a aplicação localmente, é necessário ter instalado:

1. [Go](https://go.dev/doc/install)
1. [Docker](https://www.docker.com/products/[docker-desktop/)*

*Você pode executar a aplicação usando [SQLite em memória](#usando-sqlite).

## Execução

Para rodar a aplicação, crie um arquivo `.env` na pasta raiz do projeto utilizando como referência o arquivo `.env-template`. Certifique-se de preencher corretamente as informações `DB_DRIVER` e `DB_CONNECTION` para a conexão com o banco de dados. Este projeto aceita tanto a utilização do [SQLite](#usando-sqlite) quanto do [MySQL](#usando-mysql) como banco de dados.

Após configurar corretamente a conexão ao banco de dados, para iniciar a API, utilize o seguinte comando:

```bash
$ go run cmd/main.go
Starting web server on port  :8080
```

### Usando SQLite

Para utilizar o SQLite como banco de dados da aplicação, é necessário configurar o arquivo `.env` e garantir que no arquivo `cmd/main.go`, a importação do driver esteja sendo realizada.

```txt
DB_DRIVER=sqlite3
DB_CONNECTION=:memory:
...
```

```go
import (
    ...
    // mysql
    // _ "github.com/go-sql-driver/mysql"
    // sqlite
    _ "github.com/mattn/go-sqlite3"
    ...
)
```

### Usando MySQL

Para utilizar o MySQL como banco de dados da aplicação, também é necessário configurar o arquivo `.env` e garantir que no arquivo `cmd/main.go`, a importação do driver esteja sendo realizada.

```txt
DB_DRIVER=mysql
DB_CONNECTION=root:root@tcp(mysql:3306)/gobrax
...
```

```go
import (
    ...
    // mysql
    _ "github.com/go-sql-driver/mysql"
    // sqlite
    // _ "github.com/mattn/go-sqlite3"
    ...
)
```

No entanto, `será necessário ter uma instância do MySQL em execução`. Para isso, você pode utilizar o arquivo `docker-compose.yaml` presente na raiz do projeto. Para criar um contêiner executando a instância MySQL, utilize o seguinte comando:

```bash
docker-compose up
```

## Usando a Docker Image

Para realizar o build da Docker Image, você pode usar os seguintes comandos:

```bash
docker build -t gobrax-image .
```

Este comando criará uma imagem chamada `gobrax-image` se o código e suas dependências estiverem corretos. Para criar um contêiner utilizando a imagem, utilize este comando:

```bash
docker run -p 8080:8080 gobrax-image
```
