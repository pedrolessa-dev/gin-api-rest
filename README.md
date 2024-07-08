# API REST em Go com Gin

Este é um projeto de API REST com o objetivo de gerneciar um CRUD (Create, Read, Update e Delete) de alunos, desenvolvido em Go utilizando o framework Gin. O projeto inclui integração com Docker Compose para gerenciar o PostgreSQL, documentação Swagger, testes unitários e páginas HTML.

## Estrutura do Projeto

| Diretório       | Descrição                                               |
|-----------------|---------------------------------------------------------|
| assets/         | Contém arquivos CSS para estilos da aplicação           |
| controllers/    | Controladores da aplicação                              |
| database/       | Configuração e conexão com o banco de dados             |
| docs/           | Documentação Swagger da API                             |
| models/         | Definição dos modelos de dados da aplicação             |
| routes/         | Configuração das rotas da aplicação                     |
| templates/      | Templates HTML para as páginas da aplicação             |

## Pré-requisitos

- [Git](https://git-scm.com/downloads) - Instalação do Git
- [Go](https://golang.org/doc/install) - Instalação do Go
- [Docker](https://docs.docker.com/get-docker/) - Instalação do Docker

## Executando o Projeto

1. Clone o repositório e navegue até o diretório raiz.

    ```sh
    git clone https://github.com/pedrolessa-dev/gin-api-rest.git
    cd gin-api-rest
    ```

2. Configure as variáveis de ambiente necessárias com o seguinte comando (é possível alterar os valores no arquivo `docker-compose.yml`):

    ```sh
    export PG_USER=root
    export PG_PASSWORD=postgres
    ```

3. Execute o seguinte comando para iniciar o PostgreSQL e pgAdmin via docker compose:

    ```sh
    docker-compose up
    ```

4. Execute o seguinte comando para iniciar o servidor da aplicação:

    ```sh
    go run main.go
    ```

## Acessando o Banco de Dados

Você pode acessar o banco de dados PostgreSQL utilizando o pgAdmin incluído no Docker Compose:

1. Abra o pgAdmin em <http://localhost:54321>.

2. Faça login com as credenciais (podem ser alteradas no arquivo `docker-compose.yml`):

    - Email: `seu-email@dev.com`
    - Senha: `123456`

3. Conecte-se ao servidor PostgreSQL com as seguintes informações:

    - Host: `postgres`
    - Port: `5432`
    - Username: `root`
    - Password: `postgres`
    - Database: `db_gin_api_rest`

## Acessando a Documentação

Após iniciar o servidor, é possível acessar a documentação da aplicação no Swagger através do seguinte endereço: <http://localhost:8080/swagger/index.html>

## Rodando Testes Unitários

Para rodar os testes unitários do projeto, utilize o seguinte comando no diretório raiz:

```sh
go test
```
