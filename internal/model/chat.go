package model

import "time"

type Message struct {
	ChatID     int64
	FromUserID int64
	Text       string
	SendTime   time.Time
}
