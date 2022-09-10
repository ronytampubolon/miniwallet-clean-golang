package services

import (
	"github.com/ronytampubolon/miniwallet/models"
	trx "github.com/ronytampubolon/miniwallet/models/transaction"
	"github.com/ronytampubolon/miniwallet/schemas"
)

type WalletService interface {
	InitWallet(input *schemas.InitInput) (*models.Wallet, schemas.SchemaDatabaseError)
	EnableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError)
	DisableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError)
	BalanceWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError)
	Deposit(CustomerID string, input *schemas.InputTransaction) (*trx.Deposited, schemas.SchemaDatabaseError)
	Withdrawn(CustomerID string, input *schemas.InputTransaction) (*trx.Withdrawal, schemas.SchemaDatabaseError)
}
