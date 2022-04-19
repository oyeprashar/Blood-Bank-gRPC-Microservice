package hook

import (
    "context"
    fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
)
type FindPasswordInterface interface {
	OnRequest(ctx context.Context, request *fs.FindPasswordRequest) *fs.FindPasswordResponse
	OnData(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse
	OnError(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse, err error) *fs.FindPasswordResponse
	OnResponse(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse
}

type GenericFindPasswordExecutor struct {
	FindPasswordInterface FindPasswordInterface
}

type FindPasswordController struct{
}

var FindPasswordExecutor *GenericFindPasswordExecutor

func (ge *GenericFindPasswordExecutor) OnRequest(ctx context.Context, request *fs.FindPasswordRequest) *fs.FindPasswordResponse {
	return ge.FindPasswordInterface.OnRequest(ctx,request)
}

func (ge *GenericFindPasswordExecutor) OnResponse(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	return ge.FindPasswordInterface.OnResponse(ctx,request, response)
}

func (ge *GenericFindPasswordExecutor) OnData(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	return ge.FindPasswordInterface.OnData(ctx,request, response)
}

func (ge *GenericFindPasswordExecutor) OnError(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse, err error) *fs.FindPasswordResponse {
	return ge.FindPasswordInterface.OnError(ctx,request, response, err)
}

func (rc *FindPasswordController) OnRequest(ctx context.Context, request *fs.FindPasswordRequest) *fs.FindPasswordResponse {
	return nil
}

func (rc *FindPasswordController) OnResponse(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	return nil
}

func (rc *FindPasswordController) OnData(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse) *fs.FindPasswordResponse {
	return nil
}

func (rc *FindPasswordController) OnError(ctx context.Context, request *fs.FindPasswordRequest, response *fs.FindPasswordResponse, err error) *fs.FindPasswordResponse {
	return nil
}

