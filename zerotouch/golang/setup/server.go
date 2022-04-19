package setup

import (
	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/service"
	"context"
)

type sBloodBankSystemService struct {
	fs.UnimplementedBloodBankSystemServiceServer
	
	
}

var BloodBankSystemService *sBloodBankSystemService = &sBloodBankSystemService{
    
}



func (fs *sBloodBankSystemService) ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) (*fs.FindPasswordResponse, error) {

	return service.ExecuteFindPassword(ctx, request), nil
}

func (fs *sBloodBankSystemService) ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) (*fs.AddUserResponse, error) {

	return service.ExecuteAddUser(ctx, request), nil
}

func (fs *sBloodBankSystemService) ExecuteAddUserBulk(ctx context.Context, request *fs.BulkAddUserRequest) (*fs.BulkAddUserResponse, error) {

	return service.ExecuteAddUserBulk(ctx, request), nil
}

func (fs *sBloodBankSystemService) ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) (*fs.FindBloodResponse, error) {

	return service.ExecuteFindBlood(ctx, request), nil
}

func (fs *sBloodBankSystemService) ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) (*fs.AddBloodResponse, error) {

	return service.ExecuteAddBlood(ctx, request), nil
}

func (fs *sBloodBankSystemService) ExecuteAddBloodBulk(ctx context.Context, request *fs.BulkAddBloodRequest) (*fs.BulkAddBloodResponse, error) {

	return service.ExecuteAddBloodBulk(ctx, request), nil
}



func (fs *sBloodBankSystemService) Execute(ctx context.Context, request *fs.MultiRequests) (*fs.MultiResponses, error) {

    // TO-DO
	return nil, nil
}
