package schema

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID         uint
	ItemTitle  string // Renamed from Title
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (n Note) Title() string { return n.ItemTitle } // Renamed from Title
func (n Note) Description() string { return n.CreatedAt.Format("2006-01-02") }
func (n Note) FilterValue() string { return n.ItemTitle } // Renamed from Title
