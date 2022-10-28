# API de alunos

Esse projeto é uma API REST para gerenciar alunos de uma escola

## Executar o projeto

É bem simples executar esse projeto, basta você seguir esses passos:

### Primeiro passo
  - Copiar .env.example para .env com as variáveis de ambiente, você pode usar o comando: `cp .env.example .env`

### Segundo passo
  - Fazer o build dos containers com o comando: `docker-compose up -d --build`

## Acessar a API
 - http://localhost:{APP_PORT}/

## Acessar o banco de dados
 - http://localhost:{DB_PORT}/

## Acessar o pgadmin (gerenciador do BD)
 - http://localhost:{PG_ADMIN_PORT}/

## Tecnologias
 - **[Golang](https://go.dev/)**
 - **[Gin](https://github.com/gin-gonic/gin)**
 - **[GORM](https://gorm.io/)**
 - **[Docker](https://www.docker.com/)**
 - **[PostgreSQL](https://www.postgresql.org/)**

## Arquitetura
 - REST
 - Soft delete

## Containers docker
 - golang_app
 - postgres
 - pgadmin-compose

## Endpoints
 - `GET /alunos` Retorna a lista de alunos
 - `GET /alunos/{ID}` Retorna as informações de um aluno
 - `POST /alunos` Cria um aluno, exemlo de json a ser enviado no body:
   ```json
    {
		"nome": "Daniel",
		"cpf": "23034153066",
		"rg": "140535925",
		"matricula": "120931"
    }

   ```
 - `PATCH /alunos/{ID}` Edita os dados de um aluno, exemlo de json a ser enviado no body:
    ```json
    {
		"nome": "Daniel test",
    }
   ```
   Observação: note que por ser usado o método `PATCH`, não é necessário enviar todos os dados do recurso aluno, apenas os campos que deseja editar

 - `DELETE /alunos/{ID}` Deleta um aluno 
 - `GET /alunos/busca/{CPF}` Busca as informações de um aluno pelo CPF
