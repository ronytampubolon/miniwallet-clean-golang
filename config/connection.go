package config

import (
	"github.com/ronytampubolon/miniwallet/models"
	transcation "github.com/ronytampubolon/miniwallet/models/transaction"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("wallet.db"), &gorm.Config{})
	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}
	// do database migration
	err = db.AutoMigrate(
		&models.Wallet{},
		&transcation.Deposited{},
		&transcation.Withdrawal{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
