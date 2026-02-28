package mining

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
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
	Timestamp    time.Time     // Hora da mineração
	BlockNumber  int           // Número do bloco
	Difficulty   int           // Dificuldade usada
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
	block.Timestamp = startTime
	block.Difficulty = Difficulty
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
func NewBlock(transactions []string, prevHash string, blockNumber int) Block {
	return Block{
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        0,
		MiningTime:   0,
		BlockNumber:  blockNumber,
	}
}

// Função para criar arquivo de log
func InitLogFile() error {
	file, err := os.Create("mining_log.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	header := fmt.Sprintf("=== LOG DE MINERAÇÃO - %s ===\n", time.Now().Format("02/01/2006 15:04:05"))
	header += "Data/Hora | Bloco | Dificuldade | Nonce | Tempo (s) | Hash | Transações\n"
	header += strings.Repeat("=", 100) + "\n"

	_, err = file.WriteString(header)
	return err
}

// Função para adicionar entrada no log
func LogBlock(block Block) error {
	file, err := os.OpenFile("mining_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%s | %d | %d | %d | %.2f | %s | %d\n",
		block.Timestamp.Format("02/01/2006 15:04:05"),
		block.BlockNumber,
		block.Difficulty,
		block.Nonce,
		block.MiningTime.Seconds(),
		block.Hash,
		len(block.Transactions),
	)

	_, err = file.WriteString(logEntry)
	return err
}

// Função para salvar estatísticas finais
func LogFinalStats(totalBlocks int, totalTime time.Duration, avgDifficulty float64) error {
	file, err := os.OpenFile("mining_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	footer := "\n" + strings.Repeat("=", 100) + "\n"
	footer += "ESTATÍSTICAS FINAIS:\n"
	footer += fmt.Sprintf("Total de blocos minerados: %d\n", totalBlocks)
	footer += fmt.Sprintf("Tempo total de mineração: %.2f segundos\n", totalTime.Seconds())
	footer += fmt.Sprintf("Tempo médio por bloco: %.2f segundos\n", totalTime.Seconds()/float64(totalBlocks))
	footer += fmt.Sprintf("Dificuldade média: %.1f\n", avgDifficulty)
	footer += fmt.Sprintf("Finalizado em: %s\n", time.Now().Format("02/01/2006 15:04:05"))
	footer += strings.Repeat("=", 100) + "\n"

	_, err = file.WriteString(footer)
	return err
}

func GenerateTransactionHash(transaction string) string {
    hash := sha256.Sum256([]byte(transaction))
    return fmt.Sprintf("%x", hash)
}
