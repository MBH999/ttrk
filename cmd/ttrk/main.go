package main

import (
	"fmt"
	"os"

	"github.com/MBH999/ttrk/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m, err := tui.NewModel()
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(m)

	const debugLogging = true

	if debugLogging {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
