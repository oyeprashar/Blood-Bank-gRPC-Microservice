package hook

import (
    "context"
    fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
)

type AddBloodInterface interface {
	OnRequest(ctx context.Context, request *fs.AddBloodRequest) *fs.AddBloodResponse
	OnError(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse, err error) *fs.AddBloodResponse
	OnResponse(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse) *fs.AddBloodResponse
}

type AddBloodBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddBloodRequest) *fs.BulkAddBloodResponse
	OnError(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse, err error) *fs.BulkAddBloodResponse
	OnResponse(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse) *fs.BulkAddBloodResponse
}

type GenericAddBloodExecutor struct {
	AddBloodInterface AddBloodInterface
}

type GenericAddBloodExecutorBulk struct {
	AddBloodBulkInterface AddBloodBulkInterface
}

type AddBloodController struct{
}

type BulkAddBloodController struct {
}

var AddBloodExecutor *GenericAddBloodExecutor
var BulkAddBloodExecutor *GenericAddBloodExecutorBulk

func (ge *GenericAddBloodExecutor) OnRequest(ctx context.Context, request *fs.AddBloodRequest) *fs.AddBloodResponse {
	return ge.AddBloodInterface.OnRequest(ctx,request)
}

func (ge *GenericAddBloodExecutor) OnResponse(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse) *fs.AddBloodResponse {
	return ge.AddBloodInterface.OnResponse(ctx,request, response)
}

func (ge *GenericAddBloodExecutor) OnError(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse, err error) *fs.AddBloodResponse {
	return ge.AddBloodInterface.OnError(ctx,request, response, err)
}

func (ge *GenericAddBloodExecutorBulk ) OnRequest(ctx context.Context, request *fs.BulkAddBloodRequest) *fs.BulkAddBloodResponse {
	return ge.AddBloodBulkInterface.OnRequest(ctx,request)
}

func (ge *GenericAddBloodExecutorBulk ) OnResponse(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse) *fs.BulkAddBloodResponse {
	return ge.AddBloodBulkInterface.OnResponse(ctx,request, response)
}

func (ge *GenericAddBloodExecutorBulk ) OnError(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse, err error) *fs.BulkAddBloodResponse {
	return ge.AddBloodBulkInterface.OnError(ctx,request, response, err)
}

func (rc *AddBloodController) OnRequest(ctx context.Context, request *fs.AddBloodRequest) *fs.AddBloodResponse {
	return nil
}

func (rc *AddBloodController) OnResponse(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse) *fs.AddBloodResponse {
	return nil
}

func (rc *AddBloodController) OnError(ctx context.Context, request *fs.AddBloodRequest, response *fs.AddBloodResponse, err error) *fs.AddBloodResponse {
	return nil
}

func (rc *BulkAddBloodController) OnRequest(ctx context.Context, request *fs.BulkAddBloodRequest) *fs.BulkAddBloodResponse {
	return nil
}

func (rc *BulkAddBloodController) OnResponse(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse) *fs.BulkAddBloodResponse {
	return nil
}

func (rc *BulkAddBloodController) OnError(ctx context.Context, request *fs.BulkAddBloodRequest, response *fs.BulkAddBloodResponse, err error) *fs.BulkAddBloodResponse {
	return nil
}

