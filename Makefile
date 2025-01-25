APP_NAME  = bank
BUILD_DIR = build

run:	
		@echo "Executando Aplicação..."
		@go run main.go

build:	
		@echo "Compilando binario..."
		@go build -o $(BUILD_DIR)/$(APP_NAME) .

files:	
		@find . -name '*.go' -type f

testv:
		@echo "Executando testes..."
		@go test -v -count=1 ./test/...

testt:
		@echo "Executando testes..."
		@go test -count=1 ./test/...

cover:
		@echo "Cobertura de testes..."
		@go test -coverprofile=./test/coverage/cover.out ./...
		@echo "Saida de testes..."
		@go tool cover -html=./test/coverage/cover.out -o ./test/coverage/cover.html

benchmark:
		@echo "Executando benchmark..."
		@go test -bench=.  ./test/...

clean:	
	  @echo "Limpando diretório..."
		@rm -rf $(BUILD_DIR)/*

fmt:
		@echo "Formatando o código..."
		@go fmt ./...

migrateup:
		@echo "Migrations UP..."
		@goose up

migratedown:
		@echo "Migrations DOWN..."
		@goose down

migratereset:
		@echo "Migrations DOWN..."
		@goose reset

wire:
		@echo "Google Wire Dependencies Injection..."
		@wire core/domain/account/account_di.go
		@wire core/domain/customer/customer_di.go

h:
	@echo "Comandos disponíveis:"
	@echo "  make run   	    - Executa o programa"
	@echo "  make build 	    - Compila o binário"
	@echo "  make files 	    - Lista arquivos .go"
	@echo "  make testt  	    - Executa os testes"
	@echo "  make testv  	    - Executa os testes com saida detalhada"
	@echo "  make fmt   	    - Formata o código"
	@echo "  make clean 	    - Limpa arquivos gerados"		
	@echo "  make migrateup    - Roda as Migrations"
	@echo "  make migratereset - Reseta as Migrations"
	@echo "  make migratedown  - Rollback das Migrations"
	@echo "  make wire  	    - Criar Injecao de Dependencias Wire"
	@echo "  make cover        - Relatorio de Cobertura de Testes"
	@echo "  make benchmark    - Executa Benchmarks"