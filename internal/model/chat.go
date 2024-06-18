package model

import "time"

// Message represents a message entity.
type Message struct {
	ChatID     int64
	FromUserID int64
	Text       string
	SendTime   time.Time
}
