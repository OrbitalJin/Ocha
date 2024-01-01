package notepad

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/orbitaljin/ocha/internal"
	"github.com/orbitaljin/ocha/internal/store"
	"github.com/orbitaljin/ocha/internal/store/schema"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type state int 

const (
	idle state = iota
	listing
	renaming
	creating
	editing
	viewing
)

type Notepad struct {
	db 					*store.DB
	state				state
	prestate    state
	note				schema.Note
	list 				list.Model
	ta 					textarea.Model
	view        viewport.Model
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
		prestate: idle,
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
	NewView(&m)
	return m
}

func (m Notepad) Init() tea.Cmd {
	return textarea.Blink
}

// Global view
func (m Notepad) View() string {
	switch m.state {
	case editing: return m.EditorView()
	case viewing: return m.ViewView()
	default: return m.ViewSelf()
	}
}

// Global update
func (m Notepad) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch  m.state {
	case editing: return m.UpdateEditor(msg)
	case viewing: return m.UpdateViewer(msg)
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
		case key.Matches(msg, internal.Keymap.View): m.onView()
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Change state
func (m *Notepad) setState(state state) {
	m.prestate = m.state
	m.state = state
}

// On item selection
func (m *Notepad) onSelect() {
	note := m.selectedNote()
	if note == nil {
		return
	}
	m.note = *note
	m.onEdit()
}

// On item delete
func (m *Notepad) onDelete() {
	note := m.selectedNote()
	if note == nil {
		return
	}
	m.db.DB().Delete(note)
	m.list.RemoveItem(m.list.Index())
}

// On item creation prompt
func (m *Notepad) onCreate() {
	m.setState(creating)
	m.prompt.Focus()
}

// On item rename prompt
func (m *Notepad) onRename() {
	m.setState(renaming)
	m.prompt.Focus()
}

// Switch to the editor view
func (m *Notepad) onEdit() {
	note := m.selectedNote()
	if note == nil {
		return
	}
	m.ta.SetValue(m.note.Content)
	m.setState(editing)
}

// On view
func (m *Notepad) onView() {
	switch m.state {
	case listing:
		note := m.highlightedNote()
		if note == nil {
			return
		}
		m.setState(viewing)
		m.renderView(note.Content)
	case editing:
		m.setState(viewing)
		m.renderView(m.ta.Value())
	}
}

// Get Selected Note
func (m *Notepad) selectedNote() *schema.Note {
	selected := m.list.SelectedItem()
	note, ok := selected.(schema.Note)
	switch ok {
	case true: return &note
	default: return nil
	}
}

// Get highlighted note
func (m *Notepad) highlightedNote() *schema.Note {
	current := m.list.Items()[m.list.Cursor()]
	note, ok := current.(schema.Note)
	switch ok {
	case true: return &note
	default: return nil
	}
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
	m.setState(listing)
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
	m.setState(listing)
}