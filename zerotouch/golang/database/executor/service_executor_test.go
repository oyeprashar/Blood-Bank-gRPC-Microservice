package executor_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/executor"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/mappers"
	"github.com/DATA-DOG/go-sqlmock"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
	"fmt"
	"go.uber.org/zap"
)

var Mock sqlmock.Sqlmock

func init() {
	var db *sql.DB
	var err error
	db, Mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	executor.Driver.Driver = entsql.OpenDB("mysql", db)
	viper.SetConfigType("json")
    viper.SetConfigFile("../../../../config" + "/config.json")
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("VIPER config read error",zap.Error(err))
    }
}

func TestExecuteFindPasswordServiceExecutor(t *testing.T) {
    request := &fs.FindPasswordRequest{}
    rows := sqlmock.NewRows([]string{" "}).
    AddRow(nil)
    ctx := context.Background()

    Mock.ExpectQuery("SELECT").WillReturnRows(rows)
    resp, err := executor.RequestExecutor.ExecuteFindPassword(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectQuery("SELECT").WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteFindPassword(ctx, request)
	assert.Equal(er, err)
}

func TestExecuteAddUserServiceExecutor(t *testing.T) {
    request := &fs.AddUserRequest{}
    model := mappers.MakeAddUserRequestVO(request)
	args := executor.AddUserArgs(model)
    ctx := context.Background()

    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1]).WillReturnResult(sqlmock.NewResult(1, 1))
    resp, err := executor.RequestExecutor.ExecuteAddUser(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1]).WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteAddUser(ctx, request)
    assert.Equal(er, err)
}

func TestExecuteAddUserBulkServiceExecutor(t *testing.T) {
    request := &fs.BulkAddUserRequest{}

    req := &fs.AddUserRequest{}
	request.Requests = append(request.Requests, req)
	request.Requests = append(request.Requests, req)
	var args []interface{}
	for _, r := range request.Requests {
		model := mappers.MakeAddUserRequestVO(r)
		args = append(args, executor.AddUserArgs(model)...)
	}

    ctx := context.Background()

    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3]).WillReturnResult(sqlmock.NewResult(1, 1))
    resp, err := executor.RequestExecutor.ExecuteAddUserBulk(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3]).WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteAddUserBulk(ctx, request)
    assert.Equal(er, err)
}

func TestExecuteFindBloodServiceExecutor(t *testing.T) {
    request := &fs.FindBloodRequest{}
    rows := sqlmock.NewRows([]string{" "," "," "," "," "}).
    AddRow(nil,nil,nil,nil,nil)
    ctx := context.Background()

    Mock.ExpectQuery("SELECT").WillReturnRows(rows)
    resp, err := executor.RequestExecutor.ExecuteFindBlood(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectQuery("SELECT").WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteFindBlood(ctx, request)
	assert.Equal(er, err)
}

func TestExecuteAddBloodServiceExecutor(t *testing.T) {
    request := &fs.AddBloodRequest{}
    model := mappers.MakeAddBloodRequestVO(request)
	args := executor.AddBloodArgs(model)
    ctx := context.Background()

    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3],args[4]).WillReturnResult(sqlmock.NewResult(1, 1))
    resp, err := executor.RequestExecutor.ExecuteAddBlood(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3],args[4]).WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteAddBlood(ctx, request)
    assert.Equal(er, err)
}

func TestExecuteAddBloodBulkServiceExecutor(t *testing.T) {
    request := &fs.BulkAddBloodRequest{}

    req := &fs.AddBloodRequest{}
	request.Requests = append(request.Requests, req)
	request.Requests = append(request.Requests, req)
	var args []interface{}
	for _, r := range request.Requests {
		model := mappers.MakeAddBloodRequestVO(r)
		args = append(args, executor.AddBloodArgs(model)...)
	}

    ctx := context.Background()

    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9]).WillReturnResult(sqlmock.NewResult(1, 1))
    resp, err := executor.RequestExecutor.ExecuteAddBloodBulk(ctx, request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)

    er := errors.New("Some Error")
    Mock.ExpectExec("INSERT INTO").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9]).WillReturnError(er)
    resp, err = executor.RequestExecutor.ExecuteAddBloodBulk(ctx, request)
    assert.Equal(er, err)
}


