package main

import (
	"fmt"
	"mining/internal/mining"
)

// Constante de dificuldade (quantidade de zeros exigidos no início do hash)
const Difficulty = 7

// Quantidade de blocos a serem minerados (incluindo o bloco gênesis)
const QtdBlock = 10

// log minier
const Log = false

func main() {
	// Lista de transações
	transactions := []string{
		"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
		"03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4",
		"4e07408562bedb8b60ce05c1decfe3ad16b72230995d5c92bc6572eb0aaea",
		"ef2d127de37b942a960a96987f6e9caa2e667cbbfbdd8e819e7b342ecffbcbd",
		"e1671797c52e15f763380b45e841ec32c0edg8aabbdc0f495d8f1e9e5aebec",
	}

	// Criar o bloco gênesis (primeiro bloco da cadeia)
	genesisBlock := mining.NewBlock(transactions, "0")
	fmt.Println("=== Minerando Bloco Gênesis ===")
	fmt.Printf("PrevHash: %s\n", genesisBlock.PrevHash)
	fmt.Printf("Transações: %v\n", genesisBlock.Transactions)

	// Minerar o bloco gênesis
	fmt.Println("Iniciando mineração do bloco gênesis...")
	mining.MineBlock(&genesisBlock, Difficulty, Log)

	// Exibir resultado da mineração do bloco gênesis
	fmt.Printf("Bloco Gênesis minerado!\n")
	fmt.Printf("Hash: %s\n", genesisBlock.Hash)
	fmt.Printf("Nonce: %d\n", genesisBlock.Nonce)
	fmt.Printf("Tempo gasto para minerar o bloco gênesis: %.2f segundos\n", genesisBlock.MiningTime.Seconds())

	// Validar o bloco gênesis
	if mining.ValidateBlock(genesisBlock, Difficulty) {
		fmt.Println("Bloco Gênesis válido!")
	} else {
		fmt.Println("Bloco Gênesis inválido!")
		return // Para a execução se o bloco for inválido
	}

	// Gerar uma cadeia de blocos subsequentes
	prevHash := genesisBlock.Hash
	for i := 1; i <= QtdBlock-1; i++ {
		// Criar um novo bloco com transações fictícias
		newTransactions := []string{
			fmt.Sprintf("Transação %d-1", i),
			fmt.Sprintf("Transação %d-2", i),
		}
		newBlock := mining.NewBlock(newTransactions, prevHash)
		fmt.Printf("\n=== Minerando Bloco %d ===\n", i)
		fmt.Printf("PrevHash: %s\n", newBlock.PrevHash)
		fmt.Printf("Transações: %v\n", newBlock.Transactions)

		// Minerar o novo bloco
		fmt.Printf("Iniciando mineração do bloco %d...\n", i)
		mining.MineBlock(&newBlock, Difficulty, Log)

		// Exibir resultado da mineração do novo bloco
		fmt.Printf("Bloco %d minerado!\n", i)
		fmt.Printf("Hash: %s\n", newBlock.Hash)
		fmt.Printf("Nonce: %d\n", newBlock.Nonce)
		fmt.Printf("Tempo gasto para minerar o bloco %d: %.2f segundos\n", i, newBlock.MiningTime.Seconds())

		// Validar o novo bloco
		if mining.ValidateBlock(newBlock, Difficulty) {
			fmt.Printf("Bloco %d válido!\n", i)
		} else {
			fmt.Printf("Bloco %d inválido!\n", i)
			return // Para a execução se o bloco for inválido
		}
		// Atualiza o hash anterior para o próximo bloco
		prevHash = newBlock.Hash
	}
}
