package iterator

import "fmt"

func Run() {
	b1 := Book{Name: "aaa"}
	b2 := Book{Name: "bbb"}
	b3 := Book{Name: "ccc"}
	bookShelf := BookShelf{
		Books: []Book{b1, b2, b3},
		Last:  3,
	}

	b4 := Book{Name: "ddd"}
	bs2 := BookShelf{
		Books: []Book{b1, b4},
		Last:  2,
	}
	var it BookShelfIterator = bookShelf.Iterator()

	it2 := bs2.Iterator()
	for {
		if it.HasNext() {
			fmt.Print(it.Next())
			fmt.Print("\n")
		} else {
			fmt.Print("finish!\n")
			break
		}
	}

	for {
		if it2.HasNext() {
			fmt.Print(it2.Next())
			fmt.Print("\n")
		} else {
			fmt.Print("finish!")
			break
		}
	}
}
