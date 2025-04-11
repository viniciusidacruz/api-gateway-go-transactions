package domain

type AccountRepository interface {
	Save(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	UpdateBalance(account *Account) error
}