package tui

import "github.com/charmbracelet/lipgloss"

var (
	appTitleStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("231")).Background(lipgloss.Color("63")).Bold(true).Padding(0, 2)
	subtitleStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
	containerStyle        = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63")).Padding(1, 2).MarginTop(1)
	bodyTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	sectionTitleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("213")).Bold(true)
	helpStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	emptyStateStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
	errorBannerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("231")).Background(lipgloss.Color("196")).Bold(true).Padding(0, 1).MarginTop(1)
	statusBarStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("16")).Background(lipgloss.Color("114")).Padding(0, 1).MarginTop(1)
	listTitleStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("231")).Bold(true)
	listMetaStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	inactiveCursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingRight(1)
	activeCursorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("16")).Background(lipgloss.Color("213")).Bold(true).Padding(0, 1)
	inactiveItemStyle     = lipgloss.NewStyle()
	activeItemStyle       = lipgloss.NewStyle().Background(lipgloss.Color("57")).Foreground(lipgloss.Color("231")).Bold(true)
	metricLabelStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Bold(true)
	metricValueStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("16")).Background(lipgloss.Color("186")).Bold(true).Padding(0, 1)
	timerPanelStyle       = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("63")).Padding(1, 2)
	inputPromptStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true)
	inputFieldStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")).Padding(0, 1)
	inputValueStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("231")).Bold(true)
	inputPlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
)
