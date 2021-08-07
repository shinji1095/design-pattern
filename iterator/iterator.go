package iterator

type Iterator interface {
	next() interface{}
	hasNext() bool
}
