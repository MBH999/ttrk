package main

import (
	"fmt"
	"os"

	"github.com/MBH999/ttrk/internal/homeScreen"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(homescreen.InitialModel())

	DEBUG := true

	if DEBUG {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error! %v", err)
		os.Exit(1)
	}
}

