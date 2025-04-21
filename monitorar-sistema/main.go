package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu" // Importa o pacote gopsutil para monitoramento da CPU
	"github.com/shirou/gopsutil/v3/mem" // Importa o pacote gopsutil para monitoramento de memória
)

func main() {
	for {
		// Uso da CPU
		percent, err := cpu.Percent(0, false) // Obtém o uso da CPU em porcentagem
		// O primeiro parâmetro é o tempo de espera (0 significa instantâneo)
		// O segundo parâmetro indica se deve incluir o uso de todas as CPUs

		if err != nil {
			fmt.Println("Erro ao obter uso da CPU:", err)
			continue
		}

		// Uso da Memória
		v, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Erro ao obter uso da memória:", err)
			continue
		}

		fmt.Print("\033[H\033[2J") // Limpa o terminal (para parecer dinâmico)
		fmt.Println("=== Monitor de Sistema ===")
		fmt.Printf("Uso de CPU: %.2f%%\n", percent[0])
		fmt.Printf("Memória Total: %.2f GB\n", float64(v.Total)/1e9)
		fmt.Printf("Memória Usada: %.2f GB (%.2f%%)\n", float64(v.Used)/1e9, v.UsedPercent)
		fmt.Printf("Memória Livre: %.2f GB\n", float64(v.Free)/1e9)

		time.Sleep(2 * time.Second) // Espera 2 segundos antes de atualizar
	}
}
