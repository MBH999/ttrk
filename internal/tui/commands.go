package tui

import (
	"time"

	"github.com/MBH999/ttrk/internal/storage"
	tea "github.com/charmbracelet/bubbletea"
)

type dataLoadedMsg struct {
	data storage.Data
}

type errMsg struct {
	err error
}

type saveCompleteMsg struct{}

type tickMsg time.Time

func loadDataCmd(store *storage.Store) tea.Cmd {
	return func() tea.Msg {
		data, err := store.Load()
		if err != nil {
			return errMsg{err: err}
		}
		return dataLoadedMsg{data: data}
	}
}

func saveDataCmd(store *storage.Store, data storage.Data) tea.Cmd {
	return func() tea.Msg {
		if err := store.Save(data); err != nil {
			return errMsg{err: err}
		}
		return saveCompleteMsg{}
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
