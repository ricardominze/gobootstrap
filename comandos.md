# RUN MIGRATION goose

## Instalação
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Criando um Migration
```bash
goose create add_some_column sql
# Created new file: 20170506082420_add_some_column.sql

# Ou com opção de sequência:
goose -s create add_some_column sql
# Created new file: 00001_add_some_column.sql
```

## Executando Migration
```bash
goose -dir ./migrations postgres "user=${USER} password=${PASSWORD} dbname=${DBNAME} host=${HOST} port=${PORT} sslmode=disable" up
```

## Exemplos de Conexão

### Configuração de Variáveis
```bash
# Usuário padrão
USER=""
PORT=""
HOST=""
DBNAME="bank"
PASSWORD=""
```

### SQLite
```bash
DATABASE="sqlite3"
DBSTRING="file:${DBNAME}.db" # SQLite File
DBSTRING="file::memory:?cache=shared" # SQLite Memory
```

### MySQL
```bash
USER="root"
PORT="3306"
HOST="0.0.0.0" # cat /etc/resolv.conf
DBNAME="bank"
PASSWORD="root"
DATABASE="mysql"
DBSTRING="${USER}:${PASSWORD}@tcp(${HOST}:${PORT})/${DBNAME}" # MySQL
```

### Oracle
```bash
USER="redepos"
PORT="0000"
HOST="0.0.0.0"
DBNAME="STAGE"
PASSWORD="passwd"
DATABASE="godror"
DBSTRING="${USER}/${PASSWORD}@${HOST}:${PORT}/${DBNAME}" # Oracle
```

### PostgreSQL
```bash
USER="postgres"
PORT="5432"
HOST="172.28.32.1"
DBNAME="bank"
PASSWORD="root"
DATABASE="postgres"
DBSTRING="user=${USER} password=${PASSWORD} dbname=${DBNAME} host=${HOST} port=${PORT} sslmode=disable" # PostgreSQL
