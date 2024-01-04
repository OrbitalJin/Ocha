package schema

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID         uint		   // Primary key
	ItemTitle  string    // Renamed from Title
	Content    string    // String representation of the content
	CreatedAt  time.Time // Creation timestamp
	UpdatedAt  time.Time // Lastet Editted timestamp
}

func (n Note) Title() string { return n.ItemTitle } // Renamed from Title
func (n Note) FilterValue() string { return n.ItemTitle } // Renamed from Title
func (n Note) Description() string {
	return fmt.Sprintf(
		"Created: %s Edited: %s",
		n.CreatedAt.Format("2006-01-02"),
		n.UpdatedAt.Format("2006-01-02"),
	) 
}
