package library

import "sync"

type Library struct {
	books map[string]Book
	mtx   sync.RWMutex
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}



func (l *Library) AddBook(book Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.books[book.Title]; ok {
		return ErrBookAlreadyExists
	}

	l.books[book.Title] = book

	return nil
}

func (l *Library) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	return task, nil
}

func (l *Library) ListBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book, len(l.books))

	for k, v := range l.books {
		tmp[k] = v
	}

	return tmp
}

func (l *Library) LibraryUnreadbooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	unreadBooks := make(map[string]Book)

	for title, book := range l.books {
		if !book.IsRead {
			unreadBooks[title] = book
		}
	}

	return unreadBooks
}

func (l *Library) ReadBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Read()

	l.books[title] = book

	return book, nil
}

func (l *Library) UnreadBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Unread()

	l.books[title] = book

	return book, nil
}

func (l *Library) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.books, title)

	return nil
}