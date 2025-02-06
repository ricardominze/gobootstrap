SPINNER = - \\ \| /
define spinner
  i=0; \
  chars="- \\ | /"; \
  while kill -0 $$! 2>/dev/null; do \
    i=$$(( (i+1) % 4 )); \
    printf "\r[%s] Carregando..." $$(echo $$chars | cut -d' ' -f$$((i+1))); \
    sleep 0.1; \
  done; \
  printf "\r\033[K"; \
  wait $$!; \
  printf "\r[✔] Concluído!\n";
endef

APP_NAME  = gobootstrap
BUILD_DIR = build

.PHONY: build

all:
	@echo "\e[1;34m Comandos disponíveis: \e[0m"
	@echo "  make portl   	      	  - 🚪 Lista Portas em uso"
	@echo "  make run   	          - 🏃 Executa a Aplicacao"
	@echo "  make rundock   	  - 📦 Executa Containers (Jaeger, Prometheus, Grafana)"
	@echo "  make stopdock   	  - 📦 Para Containers (Jaeger, Prometheus, Grafana)"
	@echo "  make rundockbin   	  - 📦 Executa (Container da Aplicacao) + Containers (Jaeger, Prometheus, Grafana)"
	@echo "  make stopdockbin   	  - 📦 Para (Container da Aplicacao) + Containers (Jaeger, Prometheus, Grafana)"
	@echo "  make build 	          - 🖥️  Compila o binário"
	@echo "  make files 	          - 📜 Lista arquivos \e[1;34m.go\e[0m"
	@echo "  make testt  	          - 🧪  Executa os testes"
	@echo "  make testv  	          - 📝 Executa os Testes com saida detalhada"
	@echo "  make fmt   	          - 💾 Formata o Código"
	@echo "  make clean 	          - 🌊 Limpa arquivos gerados"		
	@echo "  make migrateup    	  - 🗃️  Roda as Migrations"
	@echo "  make migratedown  	  - 🗑️  Rollback das Migrations"
	@echo "  make migratereset 	  - 🧹 Reseta as Migrations"
	@echo "  make migration <p=name> - ✏️  Cria Migration"
	@echo "  make wire  	          - 💉 \e[1;34mG\e[1;31mo\e[1;33mo\e[1;34mg\e[1;32ml\e[1;31me\e[0m \e[1;34mWire\e[0m (atualiza a injecao de dependencias)"
	@echo "  make cover        	  - ✅ Relatorio de Cobertura de Testes"
	@echo "  make benchmark    	  - 🚀 Executa Benchmarks"
	@echo "\n"

portl:
	@echo "Listando Portas..."
	@lsof -i -P -n

run:
	@echo "Executando Aplicação..."
	@go run main.go 

rundock:	
	@echo "Executando Containers(jaeger, prometheus, grafana)..."
	@docker compose -f ./docker/docker-compose-run.yaml up -d

stopdock:	
	@echo "Parando Containers(jaeger, prometheus, grafana)..."
	@docker compose -f ./docker/docker-compose-run.yaml down

rundockbin:	
	@echo "Executando (Container da Aplicacao) + Containers(jaeger, prometheus, grafana)..."
	@docker compose -f ./docker/docker-compose-build.yaml up -d

stopdockbin:	
	@echo "Parando (Container da Aplicacao) + Containers(jaeger, prometheus, grafana)..."
	@docker compose -f ./docker/docker-compose-build.yaml down

build:	
	@echo "Compilando Binário..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) . & $(call spinner)

files:	
	@find . -name '*.go' -type f & $(call spinner)

testt:
	@echo "Executando Testes..."
	@go test -count=1 ./test/... & $(call spinner)

testv:
	@echo "Executando Testes (verbose)..."
	@go test -v -count=1 ./test/... & $(call spinner)

fmt:
	@echo "Formatando o código:"
	@go fmt ./... & $(call spinner)

clean:	
	@echo "Limpando diretório..."
	@rm -rf $(BUILD_DIR)/* & $(call spinner)

migrateup:
	@echo "Migrations UP:"
	@goose up & $(call spinner)

migratedown:
	@echo "Migrations DOWN:"
	@goose down  & $(call spinner)

migratereset:
	@echo "Migrations DOWN:"
	@goose reset & $(call spinner)

migration:
	@echo "Create Migration ($(p)):"
	@goose -s create $(p) sql

wire:
	@echo "\e[1;34mG\e[1;31mo\e[1;33mo\e[1;34mg\e[1;32ml\e[1;31me\e[0m \e[1;34mWire\e[0m:"
	@wire core/domain/account/account_di.go & $(call spinner)
	@wire core/domain/customer/customer_di.go & $(call spinner)

cover:
	@echo "Cobertura de testes..."
	@go test -coverprofile=./test/coverage/cover.out ./... & $(call spinner)
	@echo "Saida de testes..."
	@go tool cover -html=./test/coverage/cover.out -o ./test/coverage/cover.html & $(call spinner)

benchmark:
	@echo "Executando benchmark..."
	@go test -bench=.  ./test/... & $(call spinner)