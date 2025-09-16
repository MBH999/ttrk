package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/MBH999/ttrk/internal/storage"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	var sections []string

	header := lipgloss.JoinVertical(lipgloss.Left,
		appTitleStyle.Render("TTRK Time Tracker"),
		subtitleStyle.Render("Track your work without leaving the keyboard."),
	)
	sections = append(sections, header)

	if m.err != nil {
		sections = append(sections, errorBannerStyle.Render(fmt.Sprintf("Error: %v", m.err)))
	}

	sections = append(sections, containerStyle.Render(m.renderContent()))

	if m.statusMessage != "" {
		sections = append(sections, statusBarStyle.Render(m.statusMessage))
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m *model) renderContent() string {
	switch m.state {
	case modeLoading:
		return bodyTextStyle.Render("Loading data...")
	case modeHome:
		items := make([]listItem, len(homeOptions))
		for i, option := range homeOptions {
			items[i] = listItem{Title: option}
		}
		sections := []string{sectionTitleStyle.Render("Select an option")}
		if len(items) > 0 {
			sections = append(sections, "", renderSelectableList(items, m.menuCursor))
		}
		sections = append(sections, "", helpStyle.Render("Use arrow keys or j/k to move, enter to select, q to quit."))
		return lipgloss.JoinVertical(lipgloss.Left, sections...)
	case modeNewProject:
		label := "Name"
		if m.inputPlaceholder != "" {
			label = fmt.Sprintf("Name (%s)", m.inputPlaceholder)
		}
		return lipgloss.JoinVertical(lipgloss.Left,
			sectionTitleStyle.Render("Create New Project"),
			"",
			metricLabelStyle.Render(label),
			renderInput(m.inputValue, m.inputPlaceholder),
			"",
			helpStyle.Render("Press enter to save, esc to cancel."),
		)
	case modeProjectList:
		items := make([]listItem, len(m.data.Projects))
		for i, project := range m.data.Projects {
			total := projectDuration(project)
			meta := fmt.Sprintf("%d tasks | %s", len(project.Tasks), formatDuration(total))
			items[i] = listItem{Title: project.Name, Meta: meta}
		}
		sections := []string{sectionTitleStyle.Render("Projects")}
		if len(items) == 0 {
			sections = append(sections, "", emptyStateStyle.Render("No projects yet. Press 'a' to add one."))
		} else {
			sections = append(sections, "", renderSelectableList(items, m.projectCursor))
		}
		sections = append(sections, "", helpStyle.Render("Enter to open, 'a' to add, esc to return, q to quit."))
		return lipgloss.JoinVertical(lipgloss.Left, sections...)
	case modeTaskList:
		if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
			return bodyTextStyle.Render("No project selected.")
		}
		project := m.data.Projects[m.activeProject]
		items := make([]listItem, len(project.Tasks))
		for i, task := range project.Tasks {
			meta := fmt.Sprintf("Total %s", formatDuration(task.TotalDuration()))
			items[i] = listItem{Title: task.Name, Meta: meta}
		}
		sections := []string{sectionTitleStyle.Render(fmt.Sprintf("Project: %s", project.Name))}
		if len(items) == 0 {
			sections = append(sections, "", emptyStateStyle.Render("No tasks yet. Press 'a' to add one."))
		} else {
			sections = append(sections, "", renderSelectableList(items, m.taskCursor))
		}
		sections = append(sections, "", helpStyle.Render("Enter to start timer, 'a' to add task, esc to go back, q to quit."))
		return lipgloss.JoinVertical(lipgloss.Left, sections...)
	case modeNewTask:
		if m.activeProject < 0 || m.activeProject >= len(m.data.Projects) {
			return bodyTextStyle.Render("Select a project first.")
		}
		project := m.data.Projects[m.activeProject]
		label := "Name"
		if m.inputPlaceholder != "" {
			label = fmt.Sprintf("Name (%s)", m.inputPlaceholder)
		}
		return lipgloss.JoinVertical(lipgloss.Left,
			sectionTitleStyle.Render(fmt.Sprintf("Add Task to %s", project.Name)),
			"",
			metricLabelStyle.Render(label),
			renderInput(m.inputValue, m.inputPlaceholder),
			"",
			helpStyle.Render("Press enter to save, esc to cancel."),
		)
	case modeTimer:
		if m.timerProjectIndex < 0 || m.timerProjectIndex >= len(m.data.Projects) {
			return bodyTextStyle.Render("Timer stopped.")
		}
		project := m.data.Projects[m.timerProjectIndex]
		if m.timerTaskIndex < 0 || m.timerTaskIndex >= len(project.Tasks) {
			return bodyTextStyle.Render("Timer stopped.")
		}
		task := project.Tasks[m.timerTaskIndex]
		session := m.timerElapsed
		total := task.TotalDuration()
		if m.timerRunning {
			total += session
		}
		rows := []string{
			listTitleStyle.Render(fmt.Sprintf("Tracking %s / %s", project.Name, task.Name)),
			"",
			renderMetricRow("Current session", formatDuration(session)),
			renderMetricRow("Recorded total", formatDuration(total)),
			"",
			helpStyle.Render("Press space or enter to stop and log, esc to cancel, q to quit."),
		}
		return timerPanelStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
	}

	return bodyTextStyle.Render("Ready.")
}

type listItem struct {
	Title string
	Meta  string
}

func renderSelectableList(items []listItem, cursor int) string {
	if len(items) == 0 {
		return ""
	}
	rendered := make([]string, len(items))
	for i, item := range items {
		var line string
		title := listTitleStyle.Render(item.Title)
		if item.Meta != "" {
			meta := listMetaStyle.Render(item.Meta)
			line = lipgloss.JoinHorizontal(lipgloss.Left, title, " ", meta)
		} else {
			line = title
		}

		prefix := inactiveCursorStyle.Render("-")
		styled := inactiveItemStyle.Render(line)
		if i == cursor {
			prefix = activeCursorStyle.Render(">")
			styled = activeItemStyle.Render(line)
		}

		rendered[i] = lipgloss.JoinHorizontal(lipgloss.Left, prefix, " ", styled)
	}
	return strings.Join(rendered, "\n")
}

func renderMetricRow(label, value string) string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		metricLabelStyle.Render(strings.ToUpper(label)),
		" ",
		metricValueStyle.Render(value),
	)
}

func projectDuration(project storage.Project) time.Duration {
	var total time.Duration
	for _, task := range project.Tasks {
		total += task.TotalDuration()
	}
	return total
}

func formatDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	seconds := int64(d / time.Second)
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
}

func renderInput(value, placeholder string) string {
	content := value
	contentStyle := inputValueStyle
	if value == "" {
		contentStyle = inputPlaceholderStyle
		if placeholder != "" {
			content = placeholder
		}
	}

	if content == "" {
		content = " "
	}

	return lipgloss.JoinHorizontal(lipgloss.Left,
		inputPromptStyle.Render(">"),
		" ",
		inputFieldStyle.Render(contentStyle.Render(content)),
	)
}
