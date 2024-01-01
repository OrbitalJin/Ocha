package notepad

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/orbitaljin/ocha/internal"
	"github.com/orbitaljin/ocha/internal/store"
	"github.com/orbitaljin/ocha/internal/store/schema"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type state int 

const (
	listing state = iota
	renaming
	creating
	editing
	viewing
)

type Notepad struct {
	db 					*store.DB
	state				state
	note				schema.Note
	list 				list.Model
	ta 					textarea.Model
	prompt			textinput.Model
	err         error
}

func New(db *store.DB, notes []schema.Note) Notepad {
	var items []list.Item
	for _, note := range notes {
		items = append(items, note)
	}
	m := Notepad {
		state: listing,
		list:  list.New(items, list.NewDefaultDelegate(), 0, 0),
		db:    db,
		err:	 nil,
	}
	m.list.Title = "Notes"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			internal.Keymap.Create,
			internal.Keymap.Rename,
			internal.Keymap.Delete,
			internal.Keymap.View,
		}
	}
	NewEditor(&m)
	NewPrompt(&m)
	return m
}

func (m Notepad) Init() tea.Cmd {
	return textarea.Blink
}

// Global view
func (m Notepad) View() string {
	switch m.state {
	case editing: return m.EditorView()
	default: return m.ViewSelf()
	}
}

// Global update
func (m Notepad) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch  m.state {
	case editing: return m.UpdateEditor(msg)
	case creating: return m.UpdatePrompt(msg)
	case renaming: return m.UpdatePrompt(msg)
	default: return m.UpdateSelf(msg)
	}
}

// Notepad view
func (m Notepad) ViewSelf() string {
	if m.state == creating || m.state == renaming {
		return docStyle.Render(
			m.list.View(), "\n",
			m.prompt.View(),
		)
	}
	return docStyle.Render(m.list.View())
}

// Notepad update event
func (m Notepad) UpdateSelf(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter: m.onSelect()
		}
		// Handle custom keybinds
		switch {
		case key.Matches(msg, internal.Keymap.Delete): m.onDelete()
		case key.Matches(msg, internal.Keymap.Create): m.onCreate()
		case key.Matches(msg, internal.Keymap.Rename): m.onRename()
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// On item selection
func (m *Notepad) onSelect() {
	selected := m.list.SelectedItem()
	note, ok := selected.(schema.Note)
	if ok {
		m.note = note
		m.onEdit()
	}
}

// On item delete
func (m *Notepad) onDelete() {
	selected := m.list.SelectedItem()
	note, ok := selected.(schema.Note)
	if ok {
		m.db.DB().Delete(&note)
		m.list.RemoveItem(m.list.Index())
	}
}

// On item creation prompt
func (m *Notepad) onCreate() {
	m.state = creating
	m.prompt.Focus()
}

// On item rename prompt
func (m *Notepad) onRename() {
	m.state = renaming
	m.prompt.Focus()
}

// Switch to the editor view
func (m *Notepad) onEdit() {
	selected := m.list.SelectedItem()
	if selected == nil {
		return
	}
	m.ta.SetValue(m.note.Content)
	m.state = editing
}

// Save current note to the database
func (m *Notepad) saveNote() {
	m.note.Content = m.ta.Value()
	m.list.SetItem(m.list.Index(), m.note)
	m.db.DB().Save(m.note)
}

// Create a new note in the database & list it
func (m *Notepad) createNote() {
	note := schema.Note {
		ItemTitle: m.prompt.Value(),
	}
	m.db.DB().Create(&note)
	m.list.InsertItem(0, note)
	m.prompt.SetValue("")
	m.state = listing
}

// Rename a note
func (m *Notepad) renameNote() {
	selected := m.list.SelectedItem()
	note, ok := selected.(schema.Note)
	if ok {
		note.ItemTitle = m.prompt.Value()
		m.list.SetItem(m.list.Index(), note)
		m.db.DB().Save(&note)
	}
	m.prompt.SetValue("")
	m.state = listing
}