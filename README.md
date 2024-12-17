# Prisma Go Example

### Este projeto é um exemplo básico de como usar o [**Prisma Client**](https://goprisma.org/)

com Go para manipulação de dados e execução de APIs.

## Pré-requisitos

Certifique-se de ter os seguintes requisitos instalados:

- **Go** (versão 1.19 ou superior)
- **Make** (para rodar os comandos de configuração)

---

## Configuração

Antes de iniciar o projeto, você pode rodar os seguintes comandos usando `make` para configurar o Prisma Client e o banco de dados:

### 1. Comandos disponíveis com `make`

- **Limpar e organizar as dependências do Go**  
  Execute o comando para rodar `make init`, que limpa e organiza as dependências do projeto:

  ```bash
  make init
  ```

- **Gerar o cliente Prisma**  
  Execute o comando para gerar o cliente Prisma:

  ```bash
  make generate
  ```

- **Sincronizar o banco de dados com o esquema (desenvolvimento)**  
  Execute o comando para sincronizar o banco de dados com o esquema:

  ```bash
  make db-push
  ```

- **Criar o schema Prisma a partir do banco existente**  
  Execute o comando para criar o schema a partir de um banco existente:

  ```bash
  make db-pull
  ```

- **Criar migração para o ambiente de desenvolvimento**  
  Execute o comando para criar migrações para desenvolvimento:

  ```bash
  make migrate-dev
  ```

- **Sincronizar o banco de dados de produção com as migrações**  
  Execute o comando para aplicar as migrações no banco de dados de produção:

  ```bash
  make migrate-deploy
  ```

- **Rodar o projeto**  
  Execute o comando para rodar o servidor Go:

  ```bash
  make run
  ```

- **Limpar arquivos gerados pelo Prisma**  
  Execute o comando para remover os arquivos gerados:
  ```bash
  make clean
  ```

### 2. Rodando o Projeto

Para iniciar o servidor Go, execute o seguinte comando com `make`:

```bash
make run
```

O servidor será iniciado na porta padrão **5000**.

---

## Endpoints da API

Base URL: `http://localhost:5000/api/v1`

### Endpoints Disponíveis:

- **GET /todos**  
  Lista todos os "todos" cadastrados.

  Exemplo de Resposta:

  ```json
  [
    {
      "id": "cm4shxg0k0002h6fh98c2dsmy",
      "title": "Miguel",
      "description": "berta",
      "createdAt": "2024-12-17T12:00:00Z",
      "updatedAt": "2024-12-17T12:00:00Z"
    }
  ]
  ```

- **GET /todos/:id**  
  Busca um "todo" específico pelo ID.

  Exemplo de Resposta:

  ```json
  {
    "id": "cm4shxg0k0002h6fh98c2dsmy",
    "title": "Miguel",
    "description": "berta",
    "createdAt": "2024-12-17T12:00:00Z",
    "updatedAt": "2024-12-17T12:00:00Z"
  }
  ```

- **POST /todos**  
  Cria um novo "todo".

  Exemplo de Payload:

  ```json
  {
    "title": "Miguel",
    "description": "berta"
  }
  ```

  Exemplo de Resposta:

  ```json
  {
    "id": "cm4shxg0k0002h6fh98c2dsmy",
    "title": "Miguel",
    "description": "berta",
    "createdAt": "2024-12-17T12:00:00Z",
    "updatedAt": "2024-12-17T12:00:00Z"
  }
  ```

---

## Limpeza

Para remover arquivos gerados pelo Prisma, execute:

```bash
make clean
```

---

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir PRs e compartilhar feedback.

---

## Licença

Este projeto está sob a licença **MIT**. Consulte o arquivo `LICENSE` para mais detalhes.

---
