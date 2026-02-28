# Mining - Simulador de Mineração de Blockchain

Um projeto simples em Go que simula o processo de mineração de blocos em uma blockchain usando o algoritmo de Proof of Work.

## 📋 Sobre o Projeto

Este projeto implementa os conceitos básicos de mineração de blockchain, incluindo:
- Criação de blocos com transações
- Algoritmo de Proof of Work
- Cálculo de hash SHA-256
- Validação de blocos
- Sistema de dificuldade configurável

## 🚀 Como Usar

### Pré-requisitos
- Go 1.24.3 ou superior

### Executando o projeto

```bash
# Clone ou navegue para o diretório do projeto
cd mining

# Execute o programa
go run main.go
```

### Saída esperada
```
Tentando nonce: 0, Hash: a1b2c3d4e5f6...
Tentando nonce: 1, Hash: f6e5d4c3b2a1...
...
Bloco minerado!
Hash: 00a1b2c3d4e5f6...
Nonce: 152
Bloco válido!
```

## 📂 Estrutura do Projeto

```
mining/
├── go.mod                      # Módulo Go
├── main.go                     # Arquivo principal
├── README.md                   # Documentação
└── internal/
    └── mining/
        └── miner.go           # Lógica de mineração
```

## 🔧 Funcionalidades

### Block (Estrutura de Bloco)
- **Transactions**: Lista de transações (hashes)
- **PrevHash**: Hash do bloco anterior
- **Nonce**: Número usado uma vez (para prova de trabalho)
- **Hash**: Hash final do bloco

### Funções Principais

#### `NewBlock(transactions []string, prevHash string) Block`
Cria um novo bloco com as transações e hash do bloco anterior.

#### `MineBlock(block *Block)`
Executa o processo de mineração usando Proof of Work até encontrar um hash válido.

#### `ValidateBlock(block Block) bool`
Valida se um bloco atende aos critérios de dificuldade.

#### `calculateHash(block Block) string`
Calcula o hash SHA-256 do bloco.

## ⚡ Configuração

### Dificuldade
A dificuldade atual está definida como `2`, o que significa que o hash deve começar com dois zeros.

```go
const Difficulty = 2  // Hash deve começar com "00"
```

Para aumentar a dificuldade, altere este valor no arquivo [internal/mining/miner.go](internal/mining/miner.go#L10).

## 📊 Exemplo de Uso

```go
package main

import (
    "fmt"
    "mining/internal/mining"
)

func main() {
    // Lista de transações
    transactions := []string{
        "transacao1_hash",
        "transacao2_hash",
        "transacao3_hash",
    }

    // Criar novo bloco
    block := mining.NewBlock(transactions, "hash_bloco_anterior")

    // Minerar bloco
    mining.MineBlock(&block)

    // Verificar resultado
    fmt.Printf("Hash: %s\n", block.Hash)
    fmt.Printf("Nonce: %d\n", block.Nonce)
}
```

## 🔍 Como Funciona a Mineração

1. **Criação do Bloco**: Um bloco é criado com transações e o hash do bloco anterior
2. **Cálculo do Hash**: O hash é calculado usando SHA-256 com base nas transações, hash anterior e nonce
3. **Verificação da Dificuldade**: Verifica se o hash começa com o número necessário de zeros
4. **Incremento do Nonce**: Se o hash não atende à dificuldade, o nonce é incrementado
5. **Repetição**: O processo continua até encontrar um hash válido

## 🛠️ Desenvolvimento

### Testando Alterações
```bash
go run main.go
```

### Modificando a Dificuldade
Edite a constante `Difficulty` em [internal/mining/miner.go](internal/mining/miner.go#L10):
- `Difficulty = 1`: Hash deve começar com "0" (mais fácil)
- `Difficulty = 3`: Hash deve começar com "000" (mais difícil)
- `Difficulty = 4`: Hash deve começar com "0000" (muito difícil)

⚠️ **Atenção**: Aumentar muito a dificuldade pode tornar a mineração extremamente lenta.

## 📝 Licença

Este projeto é apenas para fins educacionais e demonstração dos conceitos básicos de blockchain.

---

**Desenvolvido com Go** 🔗# mining
