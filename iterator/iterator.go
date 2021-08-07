package iterator

type Iterator interface {
	Next() string
	HasNext() bool
}
