package main

import (
	"fmt"
	"mining/internal/mining"
	"time"
)

// Constante de dificuldade (quantidade de zeros exigidos no início do hash)
const Difficulty = 7

// Quantidade de blocos a serem minerados (incluindo o bloco gênesis)
const QtdBlock = 10

// log minier
const Log = false

func main() {
	// Inicializar arquivo de log
	fmt.Println("Criando arquivo de log...")
	if err := mining.InitLogFile(); err != nil {
		fmt.Printf("Erro ao criar arquivo de log: %v\n", err)
		return
	}
	fmt.Println("Arquivo de log criado: mining_log.txt")

	// Variáveis para estatísticas
	var totalTime time.Duration
	totalBlocks := QtdBlock

	// Lista de transações
	transactions := []string{
		"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
		"03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4",
		"4e07408562bedb8b60ce05c1decfe3ad16b72230995d5c92bc6572eb0aaea",
		"ef2d127de37b942a960a96987f6e9caa2e667cbbfbdd8e819e7b342ecffbcbd",
		"e1671797c52e15f763380b45e841ec32c0edg8aabbdc0f495d8f1e9e5aebec",
	}

	// Criar o bloco gênesis (primeiro bloco da cadeia)
	genesisBlock := mining.NewBlock(transactions, "0", 0)
	fmt.Println("=== Minerando Bloco Gênesis ===")
	fmt.Printf("PrevHash: %s\n", genesisBlock.PrevHash)
	fmt.Printf("Transações: %v\n", genesisBlock.Transactions)

	// Minerar o bloco gênesis
	fmt.Println("Iniciando mineração do bloco gênesis...")
	mining.MineBlock(&genesisBlock, Difficulty, Log)
	totalTime += genesisBlock.MiningTime

	// Exibir resultado da mineração do bloco gênesis
	fmt.Printf("Bloco Gênesis minerado!\n")
	fmt.Printf("Hash: %s\n", genesisBlock.Hash)
	fmt.Printf("Nonce: %d\n", genesisBlock.Nonce)
	fmt.Printf("Tempo gasto para minerar o bloco gênesis: %.2f segundos\n", genesisBlock.MiningTime.Seconds())

	// Validar o bloco gênesis
	if mining.ValidateBlock(genesisBlock, Difficulty) {
		fmt.Println("Bloco Gênesis válido!")
		if err := mining.LogBlock(genesisBlock); err != nil {
			fmt.Printf("Erro ao salvar log do bloco gênesis: %v\n", err)
		}
	} else {
		fmt.Println("Bloco Gênesis inválido!")
		return
	}

	// Gerar uma cadeia de blocos subsequentes
	prevHash := genesisBlock.Hash
	for i := 1; i <= QtdBlock-1; i++ {
		// Criar um novo bloco com transações fictícias
		newTransactions := []string{
            mining.GenerateTransactionHash(fmt.Sprintf("Transação %d-1", i)),
            mining.GenerateTransactionHash(fmt.Sprintf("Transação %d-2", i)),
			mining.GenerateTransactionHash(fmt.Sprintf("Transação %d-3", i)),
            mining.GenerateTransactionHash(fmt.Sprintf("Transação %d-4", i)),
        }
		newBlock := mining.NewBlock(newTransactions, prevHash, i)
		fmt.Printf("\n=== Minerando Bloco %d ===\n", i)
		fmt.Printf("PrevHash: %s\n", newBlock.PrevHash)
		fmt.Printf("Transações: %v\n", newBlock.Transactions)

		// Minerar o novo bloco
		fmt.Printf("Iniciando mineração do bloco %d...\n", i)
		mining.MineBlock(&newBlock, Difficulty, Log)
		totalTime += newBlock.MiningTime

		// Exibir resultado da mineração do novo bloco
		fmt.Printf("Bloco %d minerado!\n", i)
		fmt.Printf("Hash: %s\n", newBlock.Hash)
		fmt.Printf("Nonce: %d\n", newBlock.Nonce)
		fmt.Printf("Tempo gasto para minerar o bloco %d: %.2f segundos\n", i, newBlock.MiningTime.Seconds())

		// Validar o novo bloco
		if mining.ValidateBlock(newBlock, Difficulty) {
			fmt.Printf("Bloco %d válido!\n", i)
			// Registrar no log
			if err := mining.LogBlock(newBlock); err != nil {
				fmt.Printf("Erro ao salvar log do bloco %d: %v\n", i, err)
			}
		} else {
			fmt.Printf("Bloco %d inválido!\n", i)
			return // Para a execução se o bloco for inválido
		}
		// Atualiza o hash anterior para o próximo bloco
		prevHash = newBlock.Hash
	}

	// Salvar estatísticas finais no log
	fmt.Printf("\n=== MINERAÇÃO CONCLUÍDA ===")
	fmt.Printf("\nTotal de blocos minerados: %d\n", totalBlocks)
	fmt.Printf("Tempo total: %.2f segundos\n", totalTime.Seconds())
	fmt.Printf("Tempo médio por bloco: %.2f segundos\n", totalTime.Seconds()/float64(totalBlocks))
	fmt.Printf("Dificuldade: %d\n", Difficulty)

	if err := mining.LogFinalStats(totalBlocks, totalTime, float64(Difficulty)); err != nil {
		fmt.Printf("Erro ao salvar estatísticas finais: %v\n", err)
	} else {
		fmt.Println("\nEstatísticas salvas no arquivo mining_log.txt")
	}
}
