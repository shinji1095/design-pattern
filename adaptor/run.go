package adaptor

func Run() {
	var p Print = PrintBanner{
		Banner{Str: "hello"},
	}
	p.PrintWeak()
	p.PrintStrong()
}
