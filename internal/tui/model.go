package tui

import (
	"time"

	"github.com/MBH999/ttrk/internal/storage"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	modeLoading mode = iota
	modeHome
	modeNewProject
	modeProjectList
	modeTaskList
	modeNewTask
	modeTimer
)

type model struct {
	store *storage.Store
	state mode

	menuCursor    int
	projectCursor int
	taskCursor    int
	activeProject int

	inputValue       string
	inputPlaceholder string

	statusMessage string
	err           error

	timerRunning      bool
	timerStart        time.Time
	timerElapsed      time.Duration
	timerProjectIndex int
	timerTaskIndex    int

	data storage.Data
}

// NewModel constructs the initial application model.
func NewModel() (*model, error) {
	store, err := storage.DefaultStore()
	if err != nil {
		return nil, err
	}

	return &model{
		store:             store,
		state:             modeLoading,
		activeProject:     -1,
		timerProjectIndex: -1,
		timerTaskIndex:    -1,
	}, nil
}

func (m *model) Init() tea.Cmd {
	return loadDataCmd(m.store)
}
