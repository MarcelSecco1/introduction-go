package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// Inicializa o aplicativo Fyne
	a := app.New()
	// Cria uma nova janela com título "Monitor de Sistema - Marcel"
	w := a.NewWindow("Monitor de Sistema - Marcel")
	// Define o tamanho da janela
	w.Resize(fyne.NewSize(300, 200))

	// Cria rótulos para exibir o uso da CPU e da memória
	// O texto inicial é "carregando..." para indicar que os dados estão sendo obtidos
	// O uso de "widget.NewLabel" cria um novo label com o texto especificado
	cpuLabel := widget.NewLabel("CPU: carregando...")
	memLabel := widget.NewLabel("Memória: carregando...")

	// O uso de "container.NewVBox" cria um novo contêiner vertical para organizar os rótulos
	content := container.NewVBox(cpuLabel, memLabel)
	w.SetContent(content)

	// Inicia uma goroutine para atualizar os rótulos a cada 2 segundos
	go func() {
		// O loop infinito é usado para atualizar os rótulos continuamente
		for {
			percent, _ := cpu.Percent(0, false) // Obtém o uso da CPU em porcentagem
			v, _ := mem.VirtualMemory()	// Obtém informações sobre a memória virtual

			cpuLabel.SetText(fmt.Sprintf("Uso de CPU: %.2f%%", percent[0])) // Define o novo texto do label com o uso da CPU
			memLabel.SetText(fmt.Sprintf("Memória Usada: %.2f GB (%.2f%%)",	// Define o novo texto do label com o uso da memória
				float64(v.Used)/1e9, v.UsedPercent))

			canvas.Refresh(cpuLabel) // Atualiza o rótulo da CPU na tela (Redesenha o label)
			canvas.Refresh(memLabel) // Atualiza o rótulo da memória na tela (tbm redesenha o label)

			time.Sleep(2 * time.Second)
		}
	}()

	w.ShowAndRun()
}
