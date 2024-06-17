package model

type AddUsersParam struct {
	ChatID    int64
	CreatorID int64
	UserIDs   []int64
}

type KickUsersParam struct {
	ChatID    int64
	CreatorID int64
	UserIDs   []int64
}
