package hook

import (
    "context"
    "go.uber.org/zap"
)

var logger = getLogger()

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Unable to initialize log at init()")
	}
	return logger
}

type ExecuteRequestController struct{
}

var ExecuteRequestExecutor *ExecuteRequestController

func init() {
	    FindPasswordExecutor = &GenericFindPasswordExecutor{
        FindPasswordInterface: &FindPasswordController{},
    }
	AddUserExecutor = &GenericAddUserExecutor{
        AddUserInterface: &AddUserController{},
    }
	BulkAddUserExecutor = &GenericAddUserExecutorBulk{
        AddUserBulkInterface: &BulkAddUserController{},
    }
    FindBloodExecutor = &GenericFindBloodExecutor{
        FindBloodInterface: &FindBloodController{},
    }
	AddBloodExecutor = &GenericAddBloodExecutor{
        AddBloodInterface: &AddBloodController{},
    }
	BulkAddBloodExecutor = &GenericAddBloodExecutorBulk{
        AddBloodBulkInterface: &BulkAddBloodController{},
    }

	ExecuteRequestExecutor = &ExecuteRequestController{}
}



func PreStartUpHook() {
    //This will run on application boot up before gRPC server starts
}

func PostStartUpHook() {
    //This will run on application boot up after gRPC server starts
}

func (rc *ExecuteRequestController) OnRequest(ctx context.Context, request interface{}) interface{} {

	return nil
}

func (rc *ExecuteRequestController) OnResponse(ctx context.Context, request interface{}, response interface{}) interface{} {

	return nil
}

func (rc *ExecuteRequestController) OnError(ctx context.Context, request interface{}, response interface{}, err error) interface{} {

	return nil
}
