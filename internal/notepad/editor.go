package notepad

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func NewEditor(m *Notepad) {
	ta := textarea.New()
	ta.SetWidth(78)
	ta.SetHeight(20)
	ta.Placeholder = "Start typing..."
	m.ta = ta
}

// Editor View
func (m Notepad) EditorView() string {
	return fmt.Sprintf(
		"%s.\n\n%s\n\n%s",
		fmt.Sprintf("Editing: %s", m.note.ItemTitle),
		m.ta.View(),
		helpStyle("\nctrl+s: save • ctrl+p: preview • ctrl+q: quit\n"),
	) + "\n\n"
}

// Editor update view
func (m Notepad) UpdateEditor(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlQ: m.setState(listing)
		case tea.KeyCtrlS: m.saveNote()
		case tea.KeyCtrlP: m.onView()
		default:
			if !m.ta.Focused() {
				m.ta.Focus()
			}
		}
	case error:
		m.err = msg
		return m, nil
	}
	m.ta, cmd = m.ta.Update(msg)
	return m, cmd
}