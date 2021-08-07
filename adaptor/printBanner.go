package adaptor

type PrintBanner struct {
	Banner
}

func (pb PrintBanner) PrintWeak() {
	pb.Banner.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() {
	pb.Banner.ShowWithStar()
}
