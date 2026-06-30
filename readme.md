# API de Prioridade de Reposição de Peças

API REST desenvolvida em Go para gerenciamento de estoque de peças automotivas e cálculo de prioridade de reposição.

O sistema permite cadastrar peças como filtros, pastilhas de freio, baterias, radiadores e outros componentes. A API calcula quais peças precisam ser repostas com base no estoque atual, consumo diário médio, estoque mínimo, tempo de entrega do fornecedor e nível de criticidade.

## Funcionalidades

* CRUD de peças automotivas
* Filtro opcional por categoria
* Paginação de peças
* Persistência com MongoDB
* Comando para popular o banco com dados de teste
* Comando para limpar os dados de teste
* Cálculo de prioridade de reposição
* Ordenação por urgência, criticidade, vendas e nome
* Testes unitários da lógica de prioridade
* Arquitetura desacoplada entre serviço, repositório e banco de dados

## Tecnologias

* Go `1.26.4`
* MongoDB Driver v2
* Chi Router
* Google UUID
* Godotenv
* MongoDB

Principais dependências:

```text
github.com/go-chi/chi/v5
github.com/google/uuid
github.com/joho/godotenv
go.mongodb.org/mongo-driver/v2
```

> Os arquivos `go.mod` e `go.sum` devem ser versionados no Git. Não é necessário editar o `go.sum` manualmente.

Após adicionar ou remover bibliotecas, execute:

```powershell
go mod tidy
```

## Arquitetura

```text
Requisição HTTP
      ↓
Controller / Router
      ↓
Handler
      ↓
Service
      ↓
Repository
      ↓
MongoDB
```

Estrutura principal do projeto:

```text
cmd/
├── api/
│   └── main.go
│
└── seed/
    ├── main.go
    └── generator.go

internal/
├── config/
├── controller/
├── dto/
├── handler/
├── repository/
│   └── mongodb/
├── service/
├── utils/
└── ...
```

## Pré-requisitos

* Go `1.26.4`
* MongoDB local ou MongoDB Atlas
* PowerShell ou terminal compatível
* Opcional: Air para hot reload

## Instalação

Clone o repositório e entre na pasta do projeto:

```powershell
git clone <URL_DO_REPOSITORIO>
cd gear-priority-api
```

Instale e organize as dependências:

```powershell
go mod tidy
```

Crie um arquivo `.env` na raiz do projeto:

```env
PORT=8080

MONGO_URI=mongodb://localhost:27017

MONGO_DATABASE=gear_priority
```

Exemplo usando MongoDB Atlas:

```env
PORT=8080

MONGO_URI=mongodb+srv://USUARIO:SENHA@cluster.mongodb.net/?retryWrites=true&w=majority

MONGO_DATABASE=gear_priority
```

> Nunca envie seu arquivo `.env` com credenciais para o GitHub.

## Executando a API

Na raiz do projeto:

```powershell
go run ./cmd/api
```

A API ficará disponível em:

```text
http://localhost:8080
```

## Executando com Air

Instale o Air uma vez:

```powershell
go install github.com/air-verse/air@latest
```

Depois execute:

```powershell
air
```

O Air recompila e reinicia a aplicação automaticamente ao salvar arquivos Go.

## Seed do Banco de Dados

O projeto possui um comando para gerar dados de teste realistas.

### Inserir dados padrão

Insere 300 peças aleatórias mais casos de borda:

```powershell
go run ./cmd/seed insert
```

### Inserir uma quantidade personalizada

Exemplo com 1000 peças:

```powershell
go run ./cmd/seed insert 1000
```

A distribuição dos dados é:

```text
20% peças críticas
40% peças com risco médio
40% peças com estoque saudável
```

Também são criados casos especiais para testar o algoritmo:

* Estoque negativo
* Média diária de vendas igual a zero
* Prazo de entrega muito longo
* Alta criticidade

### Limpar todas as peças

> Atenção: este comando remove todos os documentos da coleção `gears`.

```powershell
go run ./cmd/seed clear
```

## Modelo de Peça

Uma peça possui os seguintes campos:

```json
{
  "id": "uuid",
  "name": "Filtro de Óleo X",
  "category": "engine",
  "currentStock": 15,
  "minimumStock": 20,
  "averageDailySales": 4,
  "leadTimeInDays": 5,
  "unitCost": 18.5,
  "criticalityLevel": 3
}
```

| Campo               |    Tipo | Descrição                        |
| ------------------- | ------: | -------------------------------- |
| `id`                |    UUID | Identificador único da peça      |
| `name`              |  string | Nome da peça                     |
| `category`          |  string | Categoria da peça                |
| `currentStock`      | inteiro | Quantidade disponível em estoque |
| `minimumStock`      | inteiro | Quantidade mínima desejada       |
| `averageDailySales` |  número | Consumo médio diário             |
| `leadTimeInDays`    | inteiro | Prazo de entrega em dias         |
| `unitCost`          |  número | Custo unitário da peça           |
| `criticalityLevel`  | inteiro | Nível de criticidade da peça     |

## Endpoints

URL base:

```text
http://localhost:8080
```

---

## Criar peça

```http
POST /gears
```

Body:

```json
{
  "name": "Filtro de Óleo X",
  "category": "engine",
  "currentStock": 15,
  "minimumStock": 20,
  "averageDailySales": 4,
  "leadTimeInDays": 5,
  "unitCost": 18.5,
  "criticalityLevel": 3
}
```

Resposta: `201 Created`

```json
{
  "message": "Gear created successfully",
  "data": {
    "id": "f66337e9-4852-4d8a-9ce7-6638c7f8bc51",
    "name": "Filtro de Óleo X",
    "category": "engine",
    "currentStock": 15,
    "minimumStock": 20,
    "averageDailySales": 4,
    "leadTimeInDays": 5,
    "unitCost": 18.5,
    "criticalityLevel": 3
  }
}
```

---

## Buscar todas as peças

```http
GET /gears
```

Resposta: `200 OK`

```json
{
  "message": "Gears found successfully",
  "data": [
    {
      "id": "f66337e9-4852-4d8a-9ce7-6638c7f8bc51",
      "name": "Filtro de Óleo X",
      "category": "engine",
      "currentStock": 15,
      "minimumStock": 20,
      "averageDailySales": 4,
      "leadTimeInDays": 5,
      "unitCost": 18.5,
      "criticalityLevel": 3
    }
  ]
}
```

---

## Filtrar peças por categoria

```http
GET /gears?category=engine
```

Exemplos:

```http
GET /gears?category=engine
GET /gears?category=brakes
GET /gears?category=electrical
```

O filtro por categoria é opcional.

---

## Buscar peça por ID

```http
GET /gears/{id}
```

Exemplo:

```http
GET /gears/f66337e9-4852-4d8a-9ce7-6638c7f8bc51
```

Resposta:

```json
{
  "message": "Gear found successfully",
  "data": {
    "id": "f66337e9-4852-4d8a-9ce7-6638c7f8bc51",
    "name": "Filtro de Óleo X",
    "category": "engine",
    "currentStock": 15,
    "minimumStock": 20,
    "averageDailySales": 4,
    "leadTimeInDays": 5,
    "unitCost": 18.5,
    "criticalityLevel": 3
  }
}
```

---

## Atualizar peça

```http
PUT /gears/{id}
```

Body:

```json
{
  "name": "Filtro de Óleo Premium",
  "category": "engine",
  "currentStock": 30,
  "minimumStock": 20,
  "averageDailySales": 4,
  "leadTimeInDays": 5,
  "unitCost": 22.9,
  "criticalityLevel": 4
}
```

Resposta:

```json
{
  "message": "Gear updated successfully",
  "data": {
    "id": "f66337e9-4852-4d8a-9ce7-6638c7f8bc51",
    "name": "Filtro de Óleo Premium",
    "category": "engine",
    "currentStock": 30,
    "minimumStock": 20,
    "averageDailySales": 4,
    "leadTimeInDays": 5,
    "unitCost": 22.9,
    "criticalityLevel": 4
  }
}
```

> Atualmente o endpoint usa `PUT`, ou seja, é esperado que todos os campos editáveis sejam enviados. Um endpoint `PATCH` poderá ser adicionado futuramente para atualizações parciais.

---

## Remover peça

```http
DELETE /gears/{id}
```

Resposta:

```json
{
  "message": "Gear deleted successfully"
}
```

---

## Buscar peças paginadas

```http
GET /gears/page?page=1&limit=20
```

Também é possível filtrar por categoria:

```http
GET /gears/page?page=1&limit=20&category=engine
```

Resposta:

```json
{
  "data": [
    {
      "id": "f66337e9-4852-4d8a-9ce7-6638c7f8bc51",
      "name": "Alternador 112",
      "category": "electrical",
      "currentStock": 27,
      "minimumStock": 29,
      "averageDailySales": 9,
      "leadTimeInDays": 18,
      "unitCost": 282.19,
      "criticalityLevel": 3
    }
  ],
  "meta": {
    "page": 1,
    "limit": 20,
    "totalItems": 303,
    "totalPages": 16
  }
}
```

Parâmetros disponíveis:

| Parâmetro  | Descrição                          |
| ---------- | ---------------------------------- |
| `page`     | Página desejada                    |
| `limit`    | Quantidade de registros por página |
| `category` | Categoria opcional para filtragem  |

---

# Prioridade de Reposição

## Buscar prioridades de reposição

```http
GET /restock/priorities
```

A rota retorna apenas peças que precisam ser repostas.

Resposta:

```json
{
  "message": "Restock priorities calculated successfully",
  "data": {
    "priorities": [
      {
        "id": "uuid-1",
        "name": "Filtro de Óleo X",
        "currentStock": 15,
        "projectedStock": -5,
        "minimumStock": 20,
        "urgencyScore": 75
      }
    ]
  }
}
```

## Regras de cálculo

### Consumo esperado durante o prazo de entrega

```text
expectedConsumption = averageDailySales × leadTimeInDays
```

### Estoque projetado

```text
projectedStock = currentStock - expectedConsumption
```

### Necessidade de reposição

Uma peça precisa de reposição quando:

```text
projectedStock < minimumStock
```

### Score de urgência

```text
urgencyScore = (minimumStock - projectedStock) × criticalityLevel
```

A expressão:

```text
minimumStock - projectedStock
```

representa a falta esperada de estoque. Portanto, o `urgencyScore` já leva em consideração a quantidade faltante e a criticidade da peça.

## Ordem de prioridade

As peças são ordenadas por:

1. Maior `urgencyScore`
2. Maior `criticalityLevel`
3. Maior `averageDailySales`
4. Ordem alfabética pelo nome da peça

## Estoque projetado negativo

Um valor negativo em `projectedStock` é esperado e representa uma falta prevista antes da chegada da reposição.

Exemplo:

```text
Estoque atual: 15
Média diária de vendas: 8
Prazo de entrega: 5 dias

Consumo esperado: 40
Estoque projetado: 15 - 40 = -25
```

Neste cenário, espera-se que faltem 25 unidades antes da chegada de uma nova reposição.

## Formato de erros

A API retorna erros neste formato:

```json
{
  "error": "Invalid request body"
}
```

Outros exemplos:

```json
{
  "error": "Gear not found"
}
```

```json
{
  "error": "Invalid UUID"
}
```

## Testes

Executar todos os testes:

```powershell
go test ./...
```

Executar apenas os testes da lógica de prioridade:

```powershell
go test ./internal/service
```

Executar testes com cobertura:

```powershell
go test ./internal/service -cover
```

Os testes de prioridade verificam:

* Cálculo de consumo esperado
* Estoque projetado
* Necessidade de reposição
* Estoque negativo
* Média diária de vendas igual a zero
* Cálculo do score de urgência
* Ordenação e critérios de desempate
* Exclusão de peças com estoque saudável

## Melhorias Futuras

* Implementar `PATCH /gears/{id}` para atualizações parciais
* Validar valores negativos e limites de criticidade
* Adicionar fornecedores
* Criar pedidos de compra
* Criar autenticação e autorização
* Adicionar índices no MongoDB
* Implementar Swagger/OpenAPI
* Criar testes de integração com MongoDB
* Implementar paginação por cursor para coleções muito grandes

## Licença

Projeto criado para estudo, portfólio e experimentação de lógica de gestão de estoque automotivo.
