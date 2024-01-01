package internal

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Launch(model tea.Model) error {
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured in %s: %v", model, err)
		return err
	}
	return nil
}