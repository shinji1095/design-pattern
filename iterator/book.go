package iterator

type Book struct {
	Name string
}

func (b *Book) GetName() string {
	return b.Name
}
