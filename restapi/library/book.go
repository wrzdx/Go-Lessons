package library

import "time"

type Book struct {
	Title   string
	Author  string
	Size    int
	IsRead  bool
	AddedAt time.Time
	ReadAt  *time.Time
}

func NewBook(title string, author string, size int) Book {
	return Book{
		Title:   title,
		Author:  author,
		Size:    size,
		IsRead:  false,
		AddedAt: time.Now(),
		ReadAt:  nil,
	}
}

func (b *Book) Read() {
	readTime := time.Now()

	b.IsRead = true
	b.ReadAt = &readTime
}

func (b *Book) Unread() {
	b.IsRead = false
	b.ReadAt = nil
}