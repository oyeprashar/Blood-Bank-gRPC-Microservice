package service

import (
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/executor"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/metrics"
	"code.nurture.farm/BloodBankSystemService/core/golang/hook"
	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"context"
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger = getLogger()

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize logger, err: %v", err))
	}
	return logger
}

const (
	
	MULTI_REQUEST = "NF_RG_MULTI_REQUEST"
)

func ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) *fs.FindPasswordResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindPassword_Metrics,"FindPassword",&err,ctx)
	logger.Info("Serving FindPassword request", zap.Any("request", request))

	onRequestResponse := hook.FindPasswordExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteFindPassword request", zap.Any("request", request))
        return response
    }

	response, err := executor.RequestExecutor.ExecuteFindPassword(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindPassword request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindPassword_Error_Metrics,err,ctx)
		response = &fs.FindPasswordResponse{
             Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

        }

		onErrorResponse := hook.FindPasswordExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
            response = onErrorResponse
         }
	    return response
	}

	onDataResponse := hook.FindPasswordExecutor.OnData(ctx, request, response)
    if onDataResponse != nil {
        response := onDataResponse
        logger.Info("Returning OnData response for ExecuteeFindPassword request", zap.Any("request", request))
        return response
    }

	//On Respponse logic can be added here


	logger.Info("ExecuteFindPassword request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) *fs.AddUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddUser_Metrics,"AddUser",&err,ctx)
	logger.Info("Serving AddUser request", zap.Any("request", request))

	onRequestResponse := hook.AddUserExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteAddUser request", zap.Any("request", request))
        return response
    }

    response, err := executor.RequestExecutor.ExecuteAddUser(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddUser request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddUser_Error_Metrics,err,ctx)
		response = &fs.AddUserResponse{
             Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

        }

		onErrorResponse := hook.AddUserExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
            response = onErrorResponse
         }
	    return response
	}

	//On Respponse logic can be added here


	logger.Info("ExecuteAddUser request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddUserBulk(ctx context.Context, request *fs.BulkAddUserRequest) *fs.BulkAddUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddUser_Metrics,"AddUserBulk",&err,ctx)
	logger.Info("Serving ExecuteAddUserBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddUserExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteAddUserBulk request", zap.Any("request", request))
        return response
    }

    response, err := executor.RequestExecutor.ExecuteAddUserBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddUserBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddUser_Error_Metrics,err,ctx)
		response = &fs.BulkAddUserResponse{
			Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

		}

		onErrorResponse := hook.BulkAddUserExecutor.OnError(ctx, request, nil, err)
        if onErrorResponse != nil {
            response = onErrorResponse
        }
        return response
	}

	//On Respponse logic can be added here


	logger.Info("ExecuteAddUserBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) *fs.FindBloodResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindBlood_Metrics,"FindBlood",&err,ctx)
	logger.Info("Serving FindBlood request", zap.Any("request", request))

	onRequestResponse := hook.FindBloodExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteFindBlood request", zap.Any("request", request))
        return response
    }

	response, err := executor.RequestExecutor.ExecuteFindBlood(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindBlood request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindBlood_Error_Metrics,err,ctx)
		response = &fs.FindBloodResponse{
             Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

        }

		onErrorResponse := hook.FindBloodExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
            response = onErrorResponse
         }
	    return response
	}

	onDataResponse := hook.FindBloodExecutor.OnData(ctx, request, response)
    if onDataResponse != nil {
        response := onDataResponse
        logger.Info("Returning OnData response for ExecuteeFindBlood request", zap.Any("request", request))
        return response
    }

	//On Respponse logic can be added here


	logger.Info("ExecuteFindBlood request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) *fs.AddBloodResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddBlood_Metrics,"AddBlood",&err,ctx)
	logger.Info("Serving AddBlood request", zap.Any("request", request))

	onRequestResponse := hook.AddBloodExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteAddBlood request", zap.Any("request", request))
        return response
    }

    response, err := executor.RequestExecutor.ExecuteAddBlood(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddBlood request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddBlood_Error_Metrics,err,ctx)
		response = &fs.AddBloodResponse{
             Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

        }

		onErrorResponse := hook.AddBloodExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
            response = onErrorResponse
         }
	    return response
	}

	//On Respponse logic can be added here


	logger.Info("ExecuteAddBlood request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddBloodBulk(ctx context.Context, request *fs.BulkAddBloodRequest) *fs.BulkAddBloodResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddBlood_Metrics,"AddBloodBulk",&err,ctx)
	logger.Info("Serving ExecuteAddBloodBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddBloodExecutor.OnRequest(ctx, request)
    if onRequestResponse != nil {
        response := onRequestResponse
        logger.Info("Skipping ExecuteAddBloodBulk request", zap.Any("request", request))
        return response
    }

    response, err := executor.RequestExecutor.ExecuteAddBloodBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddBloodBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddBlood_Error_Metrics,err,ctx)
		response = &fs.BulkAddBloodResponse{
			Status: &fs.Status{
				Status: fs.StatusCode_DB_FAILURE,
			},

		}

		onErrorResponse := hook.BulkAddBloodExecutor.OnError(ctx, request, nil, err)
        if onErrorResponse != nil {
            response = onErrorResponse
        }
        return response
	}

	//On Respponse logic can be added here


	logger.Info("ExecuteAddBloodBulk request served successfully!", zap.Any("request", request))
	return response
}



func Execute(ctx context.Context, request *fs.MultiRequests) *fs.MultiResponses {

	/*var err error
	defer executor.PushToRequestMetrics()(MULTI_REQUEST,&err,ctx)
	logger.Info("Serving Execute request", zap.Any("request", request))

	response := ExecuteRequestExecutor.onRequest(ctx, request)
	if response != nil {
		err = response.(error)
	}
	if err != nil {
		logger.Error("Execute bad request", zap.Error(err))
		return &fs.MultiResponses{
			Status: &fs.Status{
				Status: fs.StatusCode_INVALID_REQUEST,
			},
		}
	}

	responses := []*fs.Response{}
	errs := executor.Execute(ctx, request)
	for _, err := range errs {
		if err != nil {
			logger.Error("Execute request failed", zap.Error(err))
			response := &fs.Response{
				Status: &fs.Status{
					Status: fs.StatusCode_DB_FAILURE,
				},
			}
			responses = append(responses, response)
		}
	}

	//OnDataLogic can be added here
	//On Respponse logic can be added here

	logger.Info("Execute request served successfully!", zap.Any("request", request))
	return &fs.MultiResponses{
		Status: &fs.Status{
			Status: fs.StatusCode_SUCCESS,
		},
		Response: responses,
	}*/
	return nil
}
