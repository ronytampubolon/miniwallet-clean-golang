package services

import (
	"github.com/google/uuid"
	"github.com/ronytampubolon/miniwallet/models"
	trx "github.com/ronytampubolon/miniwallet/models/transaction"
	repositories "github.com/ronytampubolon/miniwallet/repositories/wallet"
	"github.com/ronytampubolon/miniwallet/schemas"
)

type service struct {
	repository repositories.WalletRepository
}

func NewInitWalletService(repo *repositories.WalletRepository) *service {
	return &service{repository: *repo}
}

// init Wallet
func (s *service) InitWallet(input *schemas.InitInput) (*models.Wallet, schemas.SchemaDatabaseError) {
	return s.repository.InitWalletRepository(input)
}

// get Wallet Balance
func (s *service) BalanceWallet(CustomerId string) (*models.Wallet, schemas.SchemaDatabaseError) {
	return s.repository.GetBalance(CustomerId)
}

// enable wallet
func (s *service) EnableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError) {
	return s.repository.EnableWallet(CustomerID)
}

// disabled wallet
func (s *service) DisableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError) {
	return s.repository.DisableWallet(CustomerID)
}

// func deposit
func (s *service) Deposit(CustomerID string, input *schemas.InputTransaction) (*trx.Deposited, schemas.SchemaDatabaseError) {
	return s.repository.Deposit(CustomerID, input.Amount, uuid.MustParse(input.ReferenceID))
}

// func withdraw
func (s *service) Withdrawn(CustomerID string, input *schemas.InputTransaction) (*trx.Withdrawal, schemas.SchemaDatabaseError) {
	return s.repository.Withdraw(CustomerID, input.Amount, uuid.MustParse(input.ReferenceID))
}
