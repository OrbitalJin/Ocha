package notepad

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)	


func NewView(m *Notepad) {
	m.view = viewport.New(78, 20)
	m.view.Style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)
}

func (m *Notepad) ViewView() string {
	return m.view.View() + helpStyle("\n↑/↓: Navigate • esc: Quit\n")
}

func (m *Notepad) UpdateViewer(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc: m.setState(m.prestate)
			return m, nil
		default:
			var cmd tea.Cmd
			m.view, cmd = m.view.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m *Notepad) renderView(content string) {
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(78),
	)

	str, _ := renderer.Render(content)
	m.view.SetContent(str)
}