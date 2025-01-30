# Controle de Tráfego Aéreo ✈️

Um sistema básico para gerenciamento de voos usando Golang, Gin e PostgreSQL.

## 🚀 Tecnologias
- Golang
- Gin (Framework HTTP)
- PostgreSQL + SQLX
- JWT para autenticação

## 🛠️ Configuração
1. Instale o PostgreSQL e crie o banco:
   ```sql
   CREATE DATABASE controle_trafego;
   ```
2. Configure o banco no `config/config.go`
3. Rode a aplicação:
   ```sh
   go mod tidy
   go run cmd/main.go
   ```

## 📡 Endpoints
- `GET /voos` → Lista todos os voos

### 🔜 Melhorias Futuras:
- Adicionar autenticação JWT
- Criar frontend em React

Feito por **Gustavo Costa** 🚀