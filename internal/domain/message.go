package domain

import (
	"time"
)

type Message struct {
	Id         int64
	Text       string
	Time       time.Time
	FromUserId int64
}
