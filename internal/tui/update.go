package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/MBH999/ttrk/internal/storage"
	tea "github.com/charmbracelet/bubbletea"
)

var homeOptions = []string{"New Project", "View Projects", "Quit"}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch typed := msg.(type) {
	case errMsg:
		m.err = typed.err
		return m, nil
	case dataLoadedMsg:
		m.data = typed.data
		m.state = modeHome
		m.err = nil
		return m, nil
	case saveCompleteMsg:
		return m, nil
	case tickMsg:
		if m.timerRunning {
			elapsed := time.Since(m.timerStart).Round(time.Second)
			if elapsed < 0 {
				elapsed = 0
			}
			m.timerElapsed = elapsed
			return m, tickCmd()
		}
		return m, nil
	}

	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "ctrl+c" {
			if m.state == modeTimer && m.timerRunning {
				cmd := m.finishTimer(true)
				return m, tea.Batch(cmd, tea.Quit)
			}
			return m, tea.Quit
		}
	}

	switch m.state {
	case modeLoading:
		return m, nil
	case modeHome:
		return m.updateHome(msg)
	case modeNewProject:
		return m.updateNewProject(msg)
	case modeProjectList:
		return m.updateProjectList(msg)
	case modeTaskList:
		return m.updateTaskList(msg)
	case modeNewTask:
		return m.updateNewTask(msg)
	case modeTimer:
		return m.updateTimer(msg)
	default:
		return m, nil
	}
}

func (m *model) updateHome(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.menuCursor > 0 {
				m.menuCursor--
			}
		case "down", "j":
			if m.menuCursor < len(homeOptions)-1 {
				m.menuCursor++
			}
		case "enter":
			switch m.menuCursor {
			case 0:
				m.statusMessage = ""
				m.state = modeNewProject
				m.prepareTextInput("Project name")
				return m, nil
			case 1:
				m.state = modeProjectList
				if len(m.data.Projects) == 0 {
					m.projectCursor = 0
				} else if m.projectCursor >= len(m.data.Projects) {
					m.projectCursor = len(m.data.Projects) - 1
				}
				m.statusMessage = ""
				return m, nil
			case 2:
				return m, tea.Quit
			}
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) updateNewProject(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.captureInput(msg) {
			return m, nil
		}
		switch msg.String() {
		case "esc":
			m.state = modeHome
			m.inputValue = ""
			return m, nil
		case "enter":
			name := strings.TrimSpace(m.inputValue)
			if name == "" {
				m.statusMessage = "Project name cannot be empty"
				return m, nil
			}
			return m, m.createProject(name)
		}
	}

	return m, nil
}

func (m *model) updateProjectList(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.state = modeHome
			return m, nil
		case "up", "k":
			if m.projectCursor > 0 {
				m.projectCursor--
			}
		case "down", "j":
			if m.projectCursor < len(m.data.Projects)-1 {
				m.projectCursor++
			}
		case "enter":
			if len(m.data.Projects) == 0 {
				return m, nil
			}
			m.activeProject = m.projectCursor
			m.taskCursor = 0
			m.state = modeTaskList
			m.statusMessage = ""
			return m, nil
		case "n", "a":
			m.state = modeNewProject
			m.prepareTextInput("Project name")
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) updateTaskList(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
		m.state = modeProjectList
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.state = modeProjectList
			m.statusMessage = ""
			return m, nil
		case "up", "k":
			if m.taskCursor > 0 {
				m.taskCursor--
			}
		case "down", "j":
			if tasks := m.currentTasks(); m.taskCursor < len(tasks)-1 {
				m.taskCursor++
			}
		case "enter":
			tasks := m.currentTasks()
			if len(tasks) == 0 {
				m.statusMessage = "Add a task before starting the timer"
				return m, nil
			}
			return m, m.startTimer()
		case "a", "n":
			m.state = modeNewTask
			m.prepareTextInput("Task name")
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) updateNewTask(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.captureInput(msg) {
			return m, nil
		}
		switch msg.String() {
		case "esc":
			m.state = modeTaskList
			m.inputValue = ""
			return m, nil
		case "enter":
			name := strings.TrimSpace(m.inputValue)
			if name == "" {
				m.statusMessage = "Task name cannot be empty"
				return m, nil
			}
			return m, m.createTask(name)
		}
	}

	return m, nil
}

func (m *model) updateTimer(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "space", "enter":
			cmd := m.finishTimer(true)
			return m, cmd
		case "esc":
			cmd := m.finishTimer(false)
			return m, cmd
		case "q":
			cmd := m.finishTimer(true)
			return m, tea.Batch(cmd, tea.Quit)
		}
	}

	return m, nil
}

func (m *model) createProject(name string) tea.Cmd {
	project := storage.Project{
		ID:   storage.NewID(),
		Name: name,
	}

	m.data.Projects = append(m.data.Projects, project)
	m.projectCursor = len(m.data.Projects) - 1
	m.activeProject = m.projectCursor
	m.taskCursor = 0
	m.state = modeTaskList
	m.inputValue = ""
	m.statusMessage = fmt.Sprintf("Created project %q", name)

	return m.persistData()
}

func (m *model) createTask(name string) tea.Cmd {
	if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
		m.statusMessage = "Select a project first"
		return nil
	}

	task := storage.Task{
		ID:   storage.NewID(),
		Name: name,
	}

	project := &m.data.Projects[m.activeProject]
	project.Tasks = append(project.Tasks, task)
	m.taskCursor = len(project.Tasks) - 1
	m.state = modeTaskList
	m.inputValue = ""
	m.statusMessage = fmt.Sprintf("Created task %q", name)

	return m.persistData()
}

func (m *model) startTimer() tea.Cmd {
	if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
		return nil
	}

	project := &m.data.Projects[m.activeProject]
	if m.taskCursor < 0 || m.taskCursor >= len(project.Tasks) {
		return nil
	}

	m.timerRunning = true
	m.timerStart = time.Now()
	m.timerElapsed = 0
	m.timerProjectIndex = m.activeProject
	m.timerTaskIndex = m.taskCursor
	m.state = modeTimer
	m.statusMessage = fmt.Sprintf("Tracking %q / %q", project.Name, project.Tasks[m.taskCursor].Name)

	return tickCmd()
}

func (m *model) finishTimer(log bool) tea.Cmd {
	if m.timerProjectIndex < 0 || m.timerProjectIndex >= len(m.data.Projects) {
		m.state = modeTaskList
		m.timerRunning = false
		return nil
	}

	project := &m.data.Projects[m.timerProjectIndex]
	if m.timerTaskIndex < 0 || m.timerTaskIndex >= len(project.Tasks) {
		m.state = modeTaskList
		m.timerRunning = false
		return nil
	}

	m.timerRunning = false
	end := time.Now()
	elapsed := end.Sub(m.timerStart)
	if elapsed < 0 {
		elapsed = 0
	}
	m.timerElapsed = elapsed

	m.activeProject = m.timerProjectIndex
	m.taskCursor = m.timerTaskIndex
	m.state = modeTaskList

	if !log {
		m.statusMessage = "Timer cancelled"
		return nil
	}

	task := &project.Tasks[m.timerTaskIndex]
	duration := task.AddEntry(m.timerStart, end)
	m.statusMessage = fmt.Sprintf("Logged %s to %q / %q", formatDuration(duration), project.Name, task.Name)

	return m.persistData()
}

func (m *model) persistData() tea.Cmd {
	return saveDataCmd(m.store, m.data)
}

const inputLimit = 256

func (m *model) prepareTextInput(placeholder string) {
	m.inputPlaceholder = placeholder
	m.inputValue = ""
}

func (m *model) captureInput(msg tea.KeyMsg) bool {
	switch msg.Type {
	case tea.KeyRunes:
		for _, r := range msg.Runes {
			if r == '\n' || r == '\r' {
				continue
			}
			if len(m.inputValue) >= inputLimit {
				break
			}
			m.inputValue += string(r)
		}
		return true
	case tea.KeySpace:
		if len(m.inputValue) < inputLimit {
			m.inputValue += " "
		}
		return true
	case tea.KeyBackspace, tea.KeyCtrlH:
		if len(m.inputValue) > 0 {
			m.inputValue = m.inputValue[:len(m.inputValue)-1]
		}
		return true
	case tea.KeyCtrlU:
		if m.inputValue != "" {
			m.inputValue = ""
		}
		return true
	}

	return false
}

func (m *model) currentTasks() []storage.Task {
	if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
		return nil
	}
	return m.data.Projects[m.activeProject].Tasks
}
