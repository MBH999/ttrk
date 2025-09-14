package homescreen

import (
	tea "github.com/charmbracelet/bubbletea"
)

func InitialModel() model {
	return model{
		choices:  []string{"New Project", "List Projects"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
