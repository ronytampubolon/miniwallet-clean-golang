package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/ronytampubolon/miniwallet/models"
	trx "github.com/ronytampubolon/miniwallet/models/transaction"

	"github.com/ronytampubolon/miniwallet/schemas"
	"github.com/ronytampubolon/miniwallet/utils"
	"gorm.io/gorm"
)

type Repository interface {
	InitRepository(input *schemas.InitInput) (string, string)
}
type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(connection *gorm.DB) *WalletRepository {
	return &WalletRepository{db: connection}
}

/*
@ Init wallet account then return token for Authorization
*/
func (r *WalletRepository) InitWalletRepository(input *schemas.InitInput) (*models.Wallet, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", input.CustomerID).First(&account).Error; err != nil {
		/*
			@notfound create new one
		*/
		account := models.Wallet{
			ID:      uuid.New(),
			OwnedBy: uuid.MustParse(input.CustomerID),
			Status:  "disabled", // as default status
			Balance: 0,
		}
		r.db.Create(&account)
	}
	return &account, schemas.SchemaDatabaseError{Type: "Invalid Token", Code: 500}
}

func (r *WalletRepository) GetBalance(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", CustomerID).First(&account).Error; err != nil {
		return &account, schemas.SchemaDatabaseError{
			Type: "account not found",
			Code: 400,
		}
	}
	// check if enabled
	if account.Status == utils.DISABLED {
		return &account, schemas.SchemaDatabaseError{
			Type: "account is disabled",
			Code: 400,
		}
	}
	return &account, schemas.SchemaDatabaseError{}
}

// enable wallet
func (r *WalletRepository) EnableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", CustomerID).First(&account).Error; err != nil {
		return &account, schemas.SchemaDatabaseError{
			Type: "account not found",
			Code: 400,
		}
	}
	// check if enabled
	if account.Status == utils.ENABLED {
		return &account, schemas.SchemaDatabaseError{
			Type: "Already enabled",
			Code: 400,
		}
	}
	// else
	updateWallet := models.Wallet{
		ID:         account.ID,
		OwnedBy:    account.OwnedBy,
		Balance:    account.Balance,
		Status:     utils.ENABLED,
		EnabledAt:  time.Now().Local(),
		DisabledAt: account.DisabledAt,
	}
	r.db.Model(&account).Updates(&updateWallet)
	return &account, schemas.SchemaDatabaseError{}
}

// disabled
func (r *WalletRepository) DisableWallet(CustomerID string) (*models.Wallet, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", CustomerID).First(&account).Error; err != nil {
		return &account, schemas.SchemaDatabaseError{
			Type: "account not found",
			Code: 400,
		}
	}
	// check if enabled
	if account.Status == utils.DISABLED {
		return &account, schemas.SchemaDatabaseError{
			Type: "Already disabled",
			Code: 400,
		}
	}
	// else
	updateWallet := models.Wallet{
		ID:         account.ID,
		OwnedBy:    account.OwnedBy,
		Balance:    account.Balance,
		Status:     utils.DISABLED,
		EnabledAt:  account.EnabledAt,
		DisabledAt: time.Now().Local(),
	}
	r.db.Model(&account).Updates(&updateWallet)
	return &account, schemas.SchemaDatabaseError{}
}

// deposited money
func (r *WalletRepository) Deposit(CustomerID string, Amount float64, ReferenceID uuid.UUID) (*trx.Deposited, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", CustomerID).First(&account).Error; err != nil {
		return &trx.Deposited{}, schemas.SchemaDatabaseError{
			Type: "account not found",
			Code: 400,
		}
	}
	// define model Transction
	// store transaction to BD
	depositTrx := trx.Deposited{
		ID:          uuid.New(),
		DepositedBy: uuid.MustParse(CustomerID),
		Status:      utils.SUCCESS,
		DepositedAt: time.Now().Local(),
		Amount:      Amount,
		ReferenceID: ReferenceID,
	}
	// check if disabled
	if account.Status == utils.DISABLED {
		// update trx -> failed
		depositTrx.Status = utils.FAILED
		r.db.Create(&depositTrx)
		return &trx.Deposited{}, schemas.SchemaDatabaseError{
			Type: "account is disabled",
			Code: 400,
		}
	}
	// else
	updateWallet := models.Wallet{
		ID:        account.ID,
		OwnedBy:   account.OwnedBy,
		Balance:   (account.Balance + Amount),
		Status:    account.Status,
		EnabledAt: account.EnabledAt,
	}
	r.db.Create(&depositTrx)

	r.db.Model(&account).Updates(&updateWallet)

	return &depositTrx, schemas.SchemaDatabaseError{}
}

// withdraw money
func (r *WalletRepository) Withdraw(CustomerID string, Amount float64, ReferenceID uuid.UUID) (*trx.Withdrawal, schemas.SchemaDatabaseError) {
	var account models.Wallet
	if err := r.db.Where("owned_by = ?", CustomerID).First(&account).Error; err != nil {
		return &trx.Withdrawal{}, schemas.SchemaDatabaseError{
			Type: "account not found",
			Code: 400,
		}
	}
	// define model Transction
	// store transaction to BD
	withdrawTrx := trx.Withdrawal{
		ID:          uuid.New(),
		WithdrawnBy: uuid.MustParse(CustomerID),
		Status:      utils.SUCCESS,
		WithdrawnAt: time.Now().Local(),
		Amount:      Amount,
		ReferenceID: ReferenceID,
	}
	// check if disabled
	if account.Status == utils.DISABLED {
		// update trx -> failed
		withdrawTrx.Status = utils.FAILED
		r.db.Create(&withdrawTrx)
		return &trx.Withdrawal{}, schemas.SchemaDatabaseError{
			Type: "account is disabled",
			Code: 400,
		}
	}
	// else
	updateWallet := models.Wallet{
		ID:        account.ID,
		OwnedBy:   account.OwnedBy,
		Balance:   (account.Balance - Amount),
		Status:    account.Status,
		EnabledAt: account.EnabledAt,
	}
	r.db.Create(&withdrawTrx)

	r.db.Model(&account).Updates(&updateWallet)

	return &withdrawTrx, schemas.SchemaDatabaseError{}
}
