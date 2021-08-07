package adaptor

type PrintBanner struct {
	Str    string
	Banner Banner
}

func (pb PrintBanner) PrintWeak() {
	pb.Banner.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() {
	pb.Banner.ShowWithStar()
}
