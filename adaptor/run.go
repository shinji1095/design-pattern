package adaptor

func Run() {
	var p PrintBanner
	p.Str = "hello"
	p.PrintWeak()
	p.PrintStrong()
}
