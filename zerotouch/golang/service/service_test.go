package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"testing"
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"code.nurture.farm/BloodBankSystemService/core/golang/hook"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/executor"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/service"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/metrics"
	"fmt"
	"go.uber.org/zap"
)

func init() {
	viper.SetConfigType("json")
	viper.SetConfigFile("../../../config" + "/config.json")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("VIPER config read error",zap.Error(err))
	}
}

type ExecutorMock struct {
	mock.Mock
}

type MetricsMock struct {
	mock.Mock
}

type HookFindPasswordMock struct {
	mock.Mock
}
type HookAddUserMock struct {
	mock.Mock
}
type HookAddUserBulkMock struct {
	mock.Mock
}
type HookFindBloodMock struct {
	mock.Mock
}
type HookAddBloodMock struct {
	mock.Mock
}
type HookAddBloodBulkMock struct {
	mock.Mock
}


func (ms *MetricsMock) PushToSummarytMetrics() func(*prometheus.SummaryVec, string, *error, context.Context) {
	return func(request *prometheus.SummaryVec, methodName string, err *error, ctx context.Context) {

		return
	}
}
func (ms *MetricsMock) PushToErrorCounterMetrics() func(*prometheus.CounterVec, error, context.Context) {
	return func(request *prometheus.CounterVec, err error, ctx context.Context) {
		return
	}
}

func (rc *HookFindPasswordMock) OnRequest(ctx context.Context, request *fs.FindPasswordRequest) *fs.FindPasswordResponse {
	var FindPasswordResponse *fs.FindPasswordResponse
	args := rc.Called(ctx, request)
	mockedFindPasswordResponse := args.Get(0)
	if mockedFindPasswordResponse != nil {
		FindPasswordResponse = mockedFindPasswordResponse.(*fs.FindPasswordResponse)
	}
	return FindPasswordResponse
}

func (rc *HookFindPasswordMock) OnResponse(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	var FindPasswordResponse *fs.FindPasswordResponse
	args := rc.Called(ctx, request, response)
	mockedFindPasswordResponse := args.Get(0)
	if mockedFindPasswordResponse != nil {
		FindPasswordResponse = mockedFindPasswordResponse.(*fs.FindPasswordResponse)
	}
	return FindPasswordResponse
}

func (rc *HookFindPasswordMock) OnError(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse, err error) *fs.FindPasswordResponse {
	var FindPasswordResponse *fs.FindPasswordResponse
	args := rc.Called(ctx, request, response, err)
	mockedFindPasswordResponse := args.Get(0)
	if mockedFindPasswordResponse != nil {
		FindPasswordResponse = mockedFindPasswordResponse.(*fs.FindPasswordResponse)
	}
	return FindPasswordResponse
}

func (rc *HookFindPasswordMock) OnData(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	var FindPasswordResponse *fs.FindPasswordResponse
	args := rc.Called(ctx, request, response)
	mockedFindPasswordResponse := args.Get(0)
	if mockedFindPasswordResponse != nil {
		FindPasswordResponse = mockedFindPasswordResponse.(*fs.FindPasswordResponse)
	}
	return FindPasswordResponse
}
func (rc *HookAddUserMock) OnRequest(ctx context.Context, request *fs.AddUserRequest) *fs.AddUserResponse {
	var AddUserResponse *fs.AddUserResponse
	args := rc.Called(ctx, request)
	mockedAddUserResponse := args.Get(0)
	if mockedAddUserResponse != nil {
		AddUserResponse = mockedAddUserResponse.(*fs.AddUserResponse)
	}
	return AddUserResponse
}

func (rc *HookAddUserMock) OnResponse(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse) *fs.AddUserResponse {
	var AddUserResponse *fs.AddUserResponse
	args := rc.Called(ctx, request, response)
	mockedAddUserResponse := args.Get(0)
	if mockedAddUserResponse != nil {
		AddUserResponse = mockedAddUserResponse.(*fs.AddUserResponse)
	}
	return AddUserResponse
}

func (rc *HookAddUserMock) OnError(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse, err error) *fs.AddUserResponse {
	var AddUserResponse *fs.AddUserResponse
	args := rc.Called(ctx, request, response, err)
	mockedAddUserResponse := args.Get(0)
	if mockedAddUserResponse != nil {
		AddUserResponse = mockedAddUserResponse.(*fs.AddUserResponse)
	}
	return AddUserResponse
}
func (rc *HookAddUserBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddUserRequest) *fs.BulkAddUserResponse {
	var AddUserBulkResponse *fs.BulkAddUserResponse
	args := rc.Called(ctx, request)
	mockedAddUserBulkResponse := args.Get(0)
	if mockedAddUserBulkResponse != nil {
		AddUserBulkResponse = mockedAddUserBulkResponse.(*fs.BulkAddUserResponse)
	}
	return AddUserBulkResponse
}

func (rc *HookAddUserBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse) *fs.BulkAddUserResponse {
	var AddUserBulkResponse *fs.BulkAddUserResponse
	args := rc.Called(ctx, request, response)
	mockedAddUserBulkResponse := args.Get(0)
	if mockedAddUserBulkResponse != nil {
		AddUserBulkResponse = mockedAddUserBulkResponse.(*fs.BulkAddUserResponse)
	}
	return AddUserBulkResponse
}

func (rc *HookAddUserBulkMock) OnError(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse, err error) *fs.BulkAddUserResponse {
	var AddUserBulkResponse *fs.BulkAddUserResponse
	args := rc.Called(ctx, request, response, err)
	mockedAddUserBulkResponse := args.Get(0)
	if mockedAddUserBulkResponse != nil {
		AddUserBulkResponse = mockedAddUserBulkResponse.(*fs.BulkAddUserResponse)
	}
	return AddUserBulkResponse
}
func (rc *HookFindBloodMock) OnRequest(ctx context.Context, request *fs.FindBloodRequest) *fs.FindBloodResponse {
	var FindBloodResponse *fs.FindBloodResponse
	args := rc.Called(ctx, request)
	mockedFindBloodResponse := args.Get(0)
	if mockedFindBloodResponse != nil {
		FindBloodResponse = mockedFindBloodResponse.(*fs.FindBloodResponse)
	}
	return FindBloodResponse
}

func (rc *HookFindBloodMock) OnResponse(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	var FindBloodResponse *fs.FindBloodResponse
	args := rc.Called(ctx, request, response)
	mockedFindBloodResponse := args.Get(0)
	if mockedFindBloodResponse != nil {
		FindBloodResponse = mockedFindBloodResponse.(*fs.FindBloodResponse)
	}
	return FindBloodResponse
}

func (rc *HookFindBloodMock) OnError(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse, err error) *fs.FindBloodResponse {
	var FindBloodResponse *fs.FindBloodResponse
	args := rc.Called(ctx, request, response, err)
	mockedFindBloodResponse := args.Get(0)
	if mockedFindBloodResponse != nil {
		FindBloodResponse = mockedFindBloodResponse.(*fs.FindBloodResponse)
	}
	return FindBloodResponse
}

func (rc *HookFindBloodMock) OnData(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	var FindBloodResponse *fs.FindBloodResponse
	args := rc.Called(ctx, request, response)
	mockedFindBloodResponse := args.Get(0)
	if mockedFindBloodResponse != nil {
		FindBloodResponse = mockedFindBloodResponse.(*fs.FindBloodResponse)
	}
	return FindBloodResponse
}
func (rc *HookAddBloodMock) OnRequest(ctx context.Context, request *fs.AddBloodRequest) *fs.AddBloodResponse {
	var AddBloodResponse *fs.AddBloodResponse
	args := rc.Called(ctx, request)
	mockedAddBloodResponse := args.Get(0)
	if mockedAddBloodResponse != nil {
		AddBloodResponse = mockedAddBloodResponse.(*fs.AddBloodResponse)
	}
	return AddBloodResponse
}

func (rc *HookAddBloodMock) OnResponse(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse) *fs.AddBloodResponse {
	var AddBloodResponse *fs.AddBloodResponse
	args := rc.Called(ctx, request, response)
	mockedAddBloodResponse := args.Get(0)
	if mockedAddBloodResponse != nil {
		AddBloodResponse = mockedAddBloodResponse.(*fs.AddBloodResponse)
	}
	return AddBloodResponse
}

func (rc *HookAddBloodMock) OnError(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse, err error) *fs.AddBloodResponse {
	var AddBloodResponse *fs.AddBloodResponse
	args := rc.Called(ctx, request, response, err)
	mockedAddBloodResponse := args.Get(0)
	if mockedAddBloodResponse != nil {
		AddBloodResponse = mockedAddBloodResponse.(*fs.AddBloodResponse)
	}
	return AddBloodResponse
}
func (rc *HookAddBloodBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddBloodRequest) *fs.BulkAddBloodResponse {
	var AddBloodBulkResponse *fs.BulkAddBloodResponse
	args := rc.Called(ctx, request)
	mockedAddBloodBulkResponse := args.Get(0)
	if mockedAddBloodBulkResponse != nil {
		AddBloodBulkResponse = mockedAddBloodBulkResponse.(*fs.BulkAddBloodResponse)
	}
	return AddBloodBulkResponse
}

func (rc *HookAddBloodBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse) *fs.BulkAddBloodResponse {
	var AddBloodBulkResponse *fs.BulkAddBloodResponse
	args := rc.Called(ctx, request, response)
	mockedAddBloodBulkResponse := args.Get(0)
	if mockedAddBloodBulkResponse != nil {
		AddBloodBulkResponse = mockedAddBloodBulkResponse.(*fs.BulkAddBloodResponse)
	}
	return AddBloodBulkResponse
}

func (rc *HookAddBloodBulkMock) OnError(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse, err error) *fs.BulkAddBloodResponse {
	var AddBloodBulkResponse *fs.BulkAddBloodResponse
	args := rc.Called(ctx, request, response, err)
	mockedAddBloodBulkResponse := args.Get(0)
	if mockedAddBloodBulkResponse != nil {
		AddBloodBulkResponse = mockedAddBloodBulkResponse.(*fs.BulkAddBloodResponse)
	}
	return AddBloodBulkResponse
}


func (se *ExecutorMock) ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) (*fs.FindPasswordResponse, error) {    
    var FindPasswordResponse *fs.FindPasswordResponse
    args := se.Called(ctx,request)
    mockedFindPasswordResponse := args.Get(0)
    if mockedFindPasswordResponse!=nil{
		FindPasswordResponse = mockedFindPasswordResponse.(*fs.FindPasswordResponse)
	}
	return FindPasswordResponse,args.Error(1)
}
func (se *ExecutorMock) ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) (*fs.AddUserResponse, error) {    
    var AddUserResponse *fs.AddUserResponse
    args := se.Called(ctx,request)
    mockedAddUserResponse := args.Get(0)
    if mockedAddUserResponse!=nil{
		AddUserResponse = mockedAddUserResponse.(*fs.AddUserResponse)
	}
	return AddUserResponse,args.Error(1)
}
func (se *ExecutorMock) ExecuteAddUserBulk(ctx context.Context, request *fs.BulkAddUserRequest) (*fs.BulkAddUserResponse, error) {    
    var AddUserResponse *fs.BulkAddUserResponse
    args := se.Called(ctx,request)
    mockedAddUserResponse := args.Get(0)
    if mockedAddUserResponse!=nil{
		AddUserResponse = mockedAddUserResponse.(*fs.BulkAddUserResponse)
	}
	return AddUserResponse,args.Error(1)
}
func (se *ExecutorMock) ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) (*fs.FindBloodResponse, error) {    
    var FindBloodResponse *fs.FindBloodResponse
    args := se.Called(ctx,request)
    mockedFindBloodResponse := args.Get(0)
    if mockedFindBloodResponse!=nil{
		FindBloodResponse = mockedFindBloodResponse.(*fs.FindBloodResponse)
	}
	return FindBloodResponse,args.Error(1)
}
func (se *ExecutorMock) ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) (*fs.AddBloodResponse, error) {    
    var AddBloodResponse *fs.AddBloodResponse
    args := se.Called(ctx,request)
    mockedAddBloodResponse := args.Get(0)
    if mockedAddBloodResponse!=nil{
		AddBloodResponse = mockedAddBloodResponse.(*fs.AddBloodResponse)
	}
	return AddBloodResponse,args.Error(1)
}
func (se *ExecutorMock) ExecuteAddBloodBulk(ctx context.Context, request *fs.BulkAddBloodRequest) (*fs.BulkAddBloodResponse, error) {    
    var AddBloodResponse *fs.BulkAddBloodResponse
    args := se.Called(ctx,request)
    mockedAddBloodResponse := args.Get(0)
    if mockedAddBloodResponse!=nil{
		AddBloodResponse = mockedAddBloodResponse.(*fs.BulkAddBloodResponse)
	}
	return AddBloodResponse,args.Error(1)
}


func TestExecuteFindPassword(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookFindPasswordMock{}
	hook.FindPasswordExecutor = &hook.GenericFindPasswordExecutor{
		FindPasswordInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.FindPasswordResponse{
		Status: Status,
	}
	request := &fs.FindPasswordRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteFindPassword", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	hookMock.On("OnData", ctx, request, mockedResponse).Return(nil)
	response := service.ExecuteFindPassword(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteFindPassword", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.FindPasswordResponse)(nil), err).Return(nil)
	response = service.ExecuteFindPassword(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}

func TestExecuteAddUserBulk(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookAddUserBulkMock{}
	hook.BulkAddUserExecutor = &hook.GenericAddUserExecutorBulk{
		AddUserBulkInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.BulkAddUserResponse{
		Status: Status,
	}
	request := &fs.BulkAddUserRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteAddUserBulk", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	
	response := service.ExecuteAddUserBulk(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteAddUserBulk", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.BulkAddUserResponse)(nil), err).Return(nil)
	response = service.ExecuteAddUserBulk(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}

func TestExecuteAddUser(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookAddUserMock{}
	hook.AddUserExecutor = &hook.GenericAddUserExecutor{
		AddUserInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.AddUserResponse{
		Status: Status,
	}
	request := &fs.AddUserRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteAddUser", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	
	response := service.ExecuteAddUser(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteAddUser", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.AddUserResponse)(nil), err).Return(nil)
	response = service.ExecuteAddUser(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}

func TestExecuteFindBlood(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookFindBloodMock{}
	hook.FindBloodExecutor = &hook.GenericFindBloodExecutor{
		FindBloodInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.FindBloodResponse{
		Status: Status,
	}
	request := &fs.FindBloodRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteFindBlood", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	hookMock.On("OnData", ctx, request, mockedResponse).Return(nil)
	response := service.ExecuteFindBlood(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteFindBlood", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.FindBloodResponse)(nil), err).Return(nil)
	response = service.ExecuteFindBlood(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}

func TestExecuteAddBloodBulk(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookAddBloodBulkMock{}
	hook.BulkAddBloodExecutor = &hook.GenericAddBloodExecutorBulk{
		AddBloodBulkInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.BulkAddBloodResponse{
		Status: Status,
	}
	request := &fs.BulkAddBloodRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteAddBloodBulk", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	
	response := service.ExecuteAddBloodBulk(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteAddBloodBulk", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.BulkAddBloodResponse)(nil), err).Return(nil)
	response = service.ExecuteAddBloodBulk(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}

func TestExecuteAddBlood(t *testing.T) {
	executorMock := &ExecutorMock{}
	executor.RequestExecutor = &executor.GenericExecutor{
		ServiceExecutor: executorMock,
	}
	metricsMock := &MetricsMock{}
	metrics.Metrics = metricsMock
	hookMock := &HookAddBloodMock{}
	hook.AddBloodExecutor = &hook.GenericAddBloodExecutor{
		AddBloodInterface: hookMock,
	}
	ctx := context.Background()

	Status :=  &fs.Status{
        Status: fs.StatusCode_SUCCESS,
    }

	mockedResponse :=&fs.AddBloodResponse{
		Status: Status,
	}
	request := &fs.AddBloodRequest{}


	metricsMock.On("PushToSummarytMetrics").Return()
	metricsMock.On("IncrementCounterMetrics").Return()
	executorMock.On("ExecuteAddBlood", ctx, request).Return(mockedResponse,nil).Once()
	hookMock.On("OnRequest", ctx, request).Return(nil)
	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
	
	response := service.ExecuteAddBlood(ctx,request)
	assert := assert.New(t)
	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)


	err := errors.New("Some Error")
	mockedResponse.Status = &fs.Status{
        Status: fs.StatusCode_DB_FAILURE,
    }
	executorMock.On("ExecuteAddBlood", ctx, request).Return(nil,err).Once()
	hookMock.On("OnError", ctx, request, (*fs.AddBloodResponse)(nil), err).Return(nil)
	response = service.ExecuteAddBlood(ctx,request)
	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
}


