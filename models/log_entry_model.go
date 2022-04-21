package models

import (
	"fmt"
	"net/http"
	"time"
)

type LogEntry struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Producer  string    `json:"producer"`
	Message   string    `gorm:"type:text" json:"message"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (logEntry *LogEntry) Bind(r *http.Request) error {
	if logEntry.Producer == "" {
		return fmt.Errorf("The field 'producer' is required.")
	}

	if logEntry.Message == "" {
		return fmt.Errorf("The field 'message' is required.")
	}

	return nil
}

func (*LogEntry) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
