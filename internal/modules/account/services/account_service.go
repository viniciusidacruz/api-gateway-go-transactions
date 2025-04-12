package services

import (
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/domain"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) Create(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}

	err = s.repository.Save(account)

	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)

	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)

	err = s.repository.UpdateBalance(account)

	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)

	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)

	return &output, nil
}

func (s *AccountService) FindById(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindById(id)
	
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)

	return &output, nil
}