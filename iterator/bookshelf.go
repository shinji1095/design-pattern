package iterator

type BookShelf struct {
	Books []Book
	Last  uint
}

func (b *BookShelf) Iterator() BookShelfIterator {
	it := BookShelfIterator{
		BookShelf: b,
	}
	return it
}
