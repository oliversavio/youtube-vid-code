package service

type Quotter interface {
	GetQuote() (string, error)
}

func GetQuote(quotter Quotter) (string, error) {
	return quotter.GetQuote()
}
