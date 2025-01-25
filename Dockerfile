# Etapa 1: Build da aplicação
FROM golang:1.23 AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos para dentro do container
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte e compila a aplicação
COPY . .
RUN go build -o app .

# Etapa 2: Criando a imagem final
FROM alpine:latest

# Instala o SQLite no container
RUN apk add --no-cache sqlite-libs

# Define o diretório de trabalho no container
WORKDIR /app

# Copia o binário da aplicação Go
COPY --from=builder /app/app .

# Copia o banco de dados SQLite
COPY bank.db .

# Expõe a porta da aplicação, se necessário
EXPOSE 8000

# Comando de inicialização do container
CMD ["./app"]
