package notepad

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


func NewPrompt(m *Notepad) {
	m.prompt = textinput.New()
	m.prompt.Placeholder = "Title"
}

func (m *Notepad) UpdatePrompt(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter: m.onPromptSubmit()
		case tea.KeyEsc: m.state = listing
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.prompt, cmd = m.prompt.Update(msg)
	return m, cmd
}

func (m *Notepad) onPromptSubmit() {
	switch m.state {
	case creating: m.createNote()
	case renaming: m.renameNote()
	}
}