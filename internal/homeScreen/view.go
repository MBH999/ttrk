package homescreen

import "fmt"

func (m model) View() string {
	s := "TTRK Time Tracker\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress space to select.\n"
	s += "\nPress enter to submit.\n"
	s += "\nPress q to quit.\n"

	return s
}
