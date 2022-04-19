package hook

import (
    "context"
    fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
)

type AddUserInterface interface {
	OnRequest(ctx context.Context, request *fs.AddUserRequest) *fs.AddUserResponse
	OnError(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse, err error) *fs.AddUserResponse
	OnResponse(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse) *fs.AddUserResponse
}

type AddUserBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddUserRequest) *fs.BulkAddUserResponse
	OnError(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse, err error) *fs.BulkAddUserResponse
	OnResponse(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse) *fs.BulkAddUserResponse
}

type GenericAddUserExecutor struct {
	AddUserInterface AddUserInterface
}

type GenericAddUserExecutorBulk struct {
	AddUserBulkInterface AddUserBulkInterface
}

type AddUserController struct{
}

type BulkAddUserController struct {
}

var AddUserExecutor *GenericAddUserExecutor
var BulkAddUserExecutor *GenericAddUserExecutorBulk

func (ge *GenericAddUserExecutor) OnRequest(ctx context.Context, request *fs.AddUserRequest) *fs.AddUserResponse {
	return ge.AddUserInterface.OnRequest(ctx,request)
}

func (ge *GenericAddUserExecutor) OnResponse(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse) *fs.AddUserResponse {
	return ge.AddUserInterface.OnResponse(ctx,request, response)
}

func (ge *GenericAddUserExecutor) OnError(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse, err error) *fs.AddUserResponse {
	return ge.AddUserInterface.OnError(ctx,request, response, err)
}

func (ge *GenericAddUserExecutorBulk ) OnRequest(ctx context.Context, request *fs.BulkAddUserRequest) *fs.BulkAddUserResponse {
	return ge.AddUserBulkInterface.OnRequest(ctx,request)
}

func (ge *GenericAddUserExecutorBulk ) OnResponse(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse) *fs.BulkAddUserResponse {
	return ge.AddUserBulkInterface.OnResponse(ctx,request, response)
}

func (ge *GenericAddUserExecutorBulk ) OnError(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse, err error) *fs.BulkAddUserResponse {
	return ge.AddUserBulkInterface.OnError(ctx,request, response, err)
}

func (rc *AddUserController) OnRequest(ctx context.Context, request *fs.AddUserRequest) *fs.AddUserResponse {
	return nil
}

func (rc *AddUserController) OnResponse(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse) *fs.AddUserResponse {
	return nil
}

func (rc *AddUserController) OnError(ctx context.Context, request *fs.AddUserRequest, response *fs.AddUserResponse, err error) *fs.AddUserResponse {
	return nil
}

func (rc *BulkAddUserController) OnRequest(ctx context.Context, request *fs.BulkAddUserRequest) *fs.BulkAddUserResponse {
	return nil
}

func (rc *BulkAddUserController) OnResponse(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse) *fs.BulkAddUserResponse {
	return nil
}

func (rc *BulkAddUserController) OnError(ctx context.Context, request *fs.BulkAddUserRequest, response *fs.BulkAddUserResponse, err error) *fs.BulkAddUserResponse {
	return nil
}

