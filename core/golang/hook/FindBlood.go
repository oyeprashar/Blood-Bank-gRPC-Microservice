package hook

import (
    "context"
    fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
)
type FindBloodInterface interface {
	OnRequest(ctx context.Context, request *fs.FindBloodRequest) *fs.FindBloodResponse
	OnData(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse
	OnError(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse, err error) *fs.FindBloodResponse
	OnResponse(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse
}

type GenericFindBloodExecutor struct {
	FindBloodInterface FindBloodInterface
}

type FindBloodController struct{
}

var FindBloodExecutor *GenericFindBloodExecutor

func (ge *GenericFindBloodExecutor) OnRequest(ctx context.Context, request *fs.FindBloodRequest) *fs.FindBloodResponse {
	return ge.FindBloodInterface.OnRequest(ctx,request)
}

func (ge *GenericFindBloodExecutor) OnResponse(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	return ge.FindBloodInterface.OnResponse(ctx,request, response)
}

func (ge *GenericFindBloodExecutor) OnData(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	return ge.FindBloodInterface.OnData(ctx,request, response)
}

func (ge *GenericFindBloodExecutor) OnError(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse, err error) *fs.FindBloodResponse {
	return ge.FindBloodInterface.OnError(ctx,request, response, err)
}

func (rc *FindBloodController) OnRequest(ctx context.Context, request *fs.FindBloodRequest) *fs.FindBloodResponse {
	return nil
}

func (rc *FindBloodController) OnResponse(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	return nil
}

func (rc *FindBloodController) OnData(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse) *fs.FindBloodResponse {
	return nil
}

func (rc *FindBloodController) OnError(ctx context.Context, request *fs.FindBloodRequest, response *fs.FindBloodResponse, err error) *fs.FindBloodResponse {
	return nil
}

