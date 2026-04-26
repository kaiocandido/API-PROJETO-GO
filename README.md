# 📱 API de Rede Social

API backend desenvolvida em Go para gerenciamento de usuários e publicações (posts), simulando funcionalidades básicas de uma rede social.

---

## 🚀 Tecnologias utilizadas

* Go (Golang)
* Fiber (framework web)
* JWT (autenticação)
* UUID

---

## 📂 Estrutura do Projeto

```
.
├── login.go          # Autenticação de usuários
├── usuarios.go       # CRUD de usuários
├── publicacoes.go    # CRUD de publicações
├── rotas.go          # Definição das rotas da API
```

---

## 🔐 Autenticação

A API utiliza autenticação baseada em **JWT (JSON Web Token)**.

### Login

**POST /login**

```json
{
  "email": "usuario@email.com",
  "senha": "123456"
}
```

**Resposta:**

```json
{
  "token": "jwt_token_aqui"
}
```

---

## 👤 Rotas de Usuários

### Criar usuário

**POST /usuarios**

```json
{
  "nome": "João",
  "email": "joao@email.com",
  "senha": "123456"
}
```

---

### Buscar usuários

**GET /usuarios**

---

### Buscar usuário por ID

**GET /usuarios/:id**

---

### Atualizar usuário

**PUT /usuarios/:id**

```json
{
  "nome": "Novo Nome",
  "email": "novo@email.com"
}
```

---

### Deletar usuário

**DELETE /usuarios/:id**

---

## 📝 Rotas de Publicações

### Criar publicação

**POST /publicacoes**

```json
{
  "titulo": "Meu post",
  "conteudo": "Conteúdo do post"
}
```

---

### Listar publicações

**GET /publicacoes**

---

### Buscar publicação por ID

**GET /publicacoes/:id**

---

### Atualizar publicação

**PUT /publicacoes/:id**

```json
{
  "titulo": "Título atualizado",
  "conteudo": "Novo conteúdo"
}
```

---

### Deletar publicação

**DELETE /publicacoes/:id**

---

## ⚙️ Como rodar o projeto

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/seu-repo.git
```

2. Acesse a pasta:

```bash
cd seu-repo
```

3. Instale as dependências:

```bash
go mod tidy
```

4. Execute o projeto:

```bash
go run main.go
```

---

## 📌 Melhorias futuras

* Sistema de seguidores
* Curtidas em publicações
* Comentários
* Upload de imagens
* Paginação

---
