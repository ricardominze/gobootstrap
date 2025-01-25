# Sobre o Projeto

Este projeto foi desenvolvido com o objetivo de estudar e aplicar, de forma prática, os conceitos de Arquitetura Limpa/Arquitetura Hexagonal. Além disso, busca demonstrar diretamente no código a implementação e o uso de diversas práticas e ferramentas. A organização dos diretórios não é apresentada como uma abordagem definitiva, mas como uma sugestão que visa facilitar a compreensão e a implementação. A seguir, são detalhadas as práticas e ferramentas utilizadas no projeto.

### ⚙️ Práticas Utilizadas

- <img src="./assets/hexagonal.png" width="70"> **Arquitetura Hexagonal**  
  - Implementação dos princípios fundamentais da Arquitetura Hexagonal (Ports e Adapters).
    <br>
- <img src="./assets/goreflection.png" width="70"> **Reflection**  
  - Utilização do arquivo `util/handler_map.go` para mapear automaticamente funções de manipulação `Handlers`.  
    <br>
- <img src="./assets/makefile.png" width="50"> **Makefile**  
  - Comandos para facilitar a execução de tarefas no projeto.  

### 🛠️ Ferramentas Utilizadas 

- <img src="./assets/gogoose.png" width="35"> **Goose** (SQL Migration)
  - Criação e execução de migrações.  
    <br>
- <img src="./assets/googlewire.png" width="45"> **Wire** (Injeção de Dependências)
  - Arquivos para configuração de injeção com Google Wire.  

### 🕵️‍♂️ Observabilidade

- <img src="./assets/otelemetry.jpg" width="30"> **OpenTelemetry**  
  - Rastreamento de código com exportação via OTLP.  
- <img src="./assets/jaeger.jpeg" width="30"> **Jaeger**
  - Integração com o Jaeger para análise de rastreamento.  


### 🧪 Testes

- <img src="./assets/gotest.png" width="45"> **Testes Unitários**: Implementação de testes para validar as funcionalidades.

