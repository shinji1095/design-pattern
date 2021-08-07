package iterator

type BookShelfIterator struct {
	BookShelf *BookShelf
	Index     uint
}

func (bi *BookShelfIterator) Next() string {
	name := bi.BookShelf.Books[bi.Index].GetName()
	bi.Index += 1
	return name
}

func (bi BookShelfIterator) HasNext() bool {
	if bi.Index >= bi.BookShelf.Last {
		return false
	} else {
		return true
	}
}
