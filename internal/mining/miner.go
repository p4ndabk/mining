package mining

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Estrutura de um bloco
type Block struct {
	Transactions []string
	PrevHash     string
	Nonce        int
	Hash         string
	MiningTime   time.Duration // Tempo gasto para minerar o bloco
}

// Função para calcular o hash do bloco
func calculateHash(block Block) string {
	record := strings.Join(block.Transactions, "") + block.PrevHash + fmt.Sprintf("%d", block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// Função para minerar (encontrar o bloco válido)
func MineBlock(block *Block, Difficulty int, Log bool) {
	startTime := time.Now()
	prefix := strings.Repeat("0", Difficulty)
	for {
		block.Hash = calculateHash(*block)
		if strings.HasPrefix(block.Hash, prefix) {
			break
		}

		// Exibe o nonce e o hash atual
		if Log {
			fmt.Printf("Nonce: %d, Hash: %s\n", block.Nonce, block.Hash)
		}
		block.Nonce++
	}
	block.MiningTime = time.Since(startTime)
}

// Função para validar se o bloco atende à dificuldade
func ValidateBlock(block Block, Difficulty int) bool {
	prefix := strings.Repeat("0", Difficulty)
	return strings.HasPrefix(block.Hash, prefix) && block.Hash == calculateHash(block)
}

// Construtor de bloco
func NewBlock(transactions []string, prevHash string) Block {
	return Block{Transactions: transactions, PrevHash: prevHash, Nonce: 0, MiningTime: 0}
}
