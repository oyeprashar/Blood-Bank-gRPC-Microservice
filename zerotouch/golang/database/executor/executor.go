package executor

import (
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database"
	"fmt"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/spf13/viper"
	"sync"
	"context"
)

type DBDriver struct {
	Driver *entsql.Driver
}

var Driver *DBDriver
var once sync.Once

func (d *DBDriver) InitializeDriver() {
	once.Do(func() {
		d.Driver, _ = database.GetDBConnection()
	})
}

func (d *DBDriver) GetDriver() *entsql.Driver{
	return d.Driver
}

func (d *DBDriver) TransactionRunner(ctx context.Context, txName string, fun Transaction) (res TransactionResult, txErr error) {
	tx, err := d.Driver.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		txErr = HandleTransactionResult(tx, txErr)
	}()
	res, txErr = fun(ctx, txName, tx)
	return res, txErr
}


type TransactionResult interface{}
type Transaction func(ctx context.Context, txName string, tx dialect.Tx) (res TransactionResult, err error)

func HandleTransactionResult(tx dialect.Tx, txErr error) error {
	if txErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			txErr = rollbackErr
		}
	} else {
		commitErr := tx.Commit()
		if commitErr != nil {
		    tx.Rollback()
			txErr = commitErr
		}
	}
	return txErr
}

func init() {

	Driver = &DBDriver{}
	Driver.InitializeDriver()
	databaseName := viper.GetString("database_name")

	fmt.Println("Successfully connected to DataBase: ",databaseName)
}
