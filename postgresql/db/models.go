package db

import "time"

type BookModel struct {
	ID      int
	Title   string
	Author  string
	Review  string
	Year    *int
	IsRead  bool
	AddedAt time.Time
	ReadAt  *time.Time
}
