package adaptor

import "fmt"

type Banner struct {
	Str string
}

func (b Banner) ShowWithParen() {
	fmt.Print("(" + b.Str + ")\n")
}

func (b Banner) ShowWithStar() {
	fmt.Print("*" + b.Str + "*\n")
}
