# Auth System

Este projeto implementa um sistema de autenticação e gestão de dívidas utilizando **Go**, com foco em boas práticas de arquitetura, testabilidade e organização de código.

## 🧠 Motivações Arquiteturais

- **Camadas bem definidas**: O projeto segue uma separação clara em camadas (`domain`, `usecases`, `repositories`, `controllers`, `dto`), facilitando a testabilidade, manutenção e escalabilidade do sistema. Procurei manter o projeto desacoplado na medida do possível, negociando entrega, qualidade e levando em consideração o escopo.
- **GORM como ORM**: Usado para abstração do banco de dados relacional, permitindo um mapeamento mais produtivo entre structs e tabelas.
- **Fiber**: Framework HTTP leve e performático inspirado no Express.js.
- **Redis**: Utilizado para armazenar tokens de sessão de forma rápida e temporária.
- **JWT**: Garante autenticação segura e stateless, com expiração.
- **Validações personalizadas**: Inclui verificação de CPF, datas no formato brasileiro, e valores de dívida. Também há espaço para o desenvolvimento de mais validações no futuro.
- **Migrations customizadas com GORM**: Inclui inserção automática de dados na inicialização da aplicação.

## 🧱 Tecnologias Utilizadas

- Go 1.23
- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- Redis
- PostgreSQL
- JWT (github.com/golang-jwt/jwt)
- Docker e Docker Compose

---

## 🚀 Como Rodar Localmente

Este projeto utiliza **Docker** e **Docker Compose** para facilitar a execução da aplicação em qualquer ambiente.

### ✅ Pré-requisitos

- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)

### 1. Clone o projeto

```bash
git clone https://github.com/mvitoriapereirac/auth-system.git
cd auth-system
```

### 2. Suba os containers

```bash
docker-compose up --build
```

### 3. Acesse a aplicação pelo Postman ou Insomnia
A API estará disponível em:

```bash
http://localhost:8080
```

## 📡 Endpoints da API
    Content-type => Todos os endpoints precisam ter no header uma chave Content-Type com o valor `application/json`
### 1. 🔐 Autenticação

#### `POST /login`

Autentica o usuário e retorna um token de sessão.

- **Body (JSON):**
```json
{
  "cpf": "12345678900",
  "senha": "sua_senha"
}
```
**Autenticação:** Todos os endpoints protegidos requerem o envio do token JWT no header `Authorization`, com o formato:  
 `Bearer <seu_token>`


### 2. Registro
        Para facilitar o desenvolvimento, aqui assumimos que todo usuário normal é um consumidor, e todo usuário admin é um usuário relacionado a uma empresa parceira.
#### `POST /register`
- **Body (JSON):**
```json
{
    "cpf" : "13285972445",
    "tipo": "C", //[opcional] pode ser C para usuários normais (consumidores) ou E para usuários de empresas parceiras
    "dataNascimento": "05-07-1999",
    "email": "mvitoriapereirac+2@br.experian.com", //Dispensa o envio do tipo. Caso possua a terminação como no exemplo, é considerado usuário de empresa parceira
    "senha": "123456"
}
```

### 3. Cadastro de dívidas
#### `POST /user/dividas`
    Apenas para usuários de empresas parceiras. Em um primeiro momento, apenas permite cadastro de dívidas para usuários normais já cadastrados na plataforma.
```json
{
    "cpfConsumidor": "13285979425",
    "divida": {
        "vencimento": "09-08-2024",
        "valor": 20.60
    }
}
```

### 4. Listagem de dívidas
#### `GET /user/dividas`
    Apenas requer o Bearer token. Apenas permitido para autores das dívidas. 

### 5. Consulta de score
#### `GET /user/score`
    Apenas requer o Bearer token. Apenas permitido para usuários normais. 

### 5. Logout
#### `POST /user/logout`
    Apenas requer o Bearer token. Apenas permitido para usuários normais. 