package model

import "time"

type Message struct {
	ChatID     int64
	FromUserID int64
	Text       string
	SendTime   time.Time
}

type ChatInfo struct {
	CreatorUserID int64
	ChatID        int64
	ChatName      string
}

type AddUsersParam struct {
	ChatInfo ChatInfo
	UserIDs  []int64
}

type KickUsersParam struct {
	ChatInfo ChatInfo
	UserIDs  []int64
}
