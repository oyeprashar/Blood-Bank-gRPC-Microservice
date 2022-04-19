package executor

import (
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/mappers"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/models"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database"
	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"context"
	"database/sql"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
)

type ServiceExecutor interface {
    ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) (*fs.FindPasswordResponse, error)
    ExecuteAddUserBulk(ctx context.Context, bulkrequest *fs.BulkAddUserRequest) (*fs.BulkAddUserResponse, error)
    ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) (*fs.AddUserResponse, error)
    ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) (*fs.FindBloodResponse, error)
    ExecuteAddBloodBulk(ctx context.Context, bulkrequest *fs.BulkAddBloodRequest) (*fs.BulkAddBloodResponse, error)
    ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) (*fs.AddBloodResponse, error)

}

type GenericExecutor struct {
	ServiceExecutor ServiceExecutor
}

type Executor struct {
}

var RequestExecutor *GenericExecutor

func (se *GenericExecutor) ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) (*fs.FindPasswordResponse, error) {
    return se.ServiceExecutor.ExecuteFindPassword(ctx, request)
}

func (se *GenericExecutor) ExecuteAddUserBulk(ctx context.Context, bulkrequest *fs.BulkAddUserRequest) (*fs.BulkAddUserResponse, error) {
    return se.ServiceExecutor.ExecuteAddUserBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) (*fs.AddUserResponse, error) {
    return se.ServiceExecutor.ExecuteAddUser(ctx, request)
}

func (se *GenericExecutor) ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) (*fs.FindBloodResponse, error) {
    return se.ServiceExecutor.ExecuteFindBlood(ctx, request)
}

func (se *GenericExecutor) ExecuteAddBloodBulk(ctx context.Context, bulkrequest *fs.BulkAddBloodRequest) (*fs.BulkAddBloodResponse, error) {
    return se.ServiceExecutor.ExecuteAddBloodBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) (*fs.AddBloodResponse, error) {
    return se.ServiceExecutor.ExecuteAddBlood(ctx, request)
}



func init() {
	RequestExecutor = &GenericExecutor{
		ServiceExecutor: &Executor{},
	}
}

func (se *Executor) ExecuteFindPassword(ctx context.Context, request *fs.FindPasswordRequest) (*fs.FindPasswordResponse , error) {

	response := &fs.FindPasswordResponse{}
	var rows = entsql.Rows{}
	model := mappers.MakeFindPasswordRequestVO(request)
    args := FindPasswordArgsReq(model)
	currQuery := database.QUERY_FINDPASSWORD
	
	err := Driver.GetDriver().Query(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindPasswordRequest", zap.Error(err))
        return nil, err
	}
	for rows.Next() {
		model := models.FindPasswordResponseVO{}
		err := rows.Scan(&model.Password)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindPasswordRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindPasswordResponseVO(&model))
	}
	response.Status = &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        }
	return response, nil
}

func (se *Executor) ExecuteAddUser(ctx context.Context, request *fs.AddUserRequest) (*fs.AddUserResponse, error) {

	model := mappers.MakeAddUserRequestVO(request)
	args := AddUserArgs(model)

	var rows sql.Result
	currQuery := database.QUERY_ADDUSER
    
	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
    if err != nil {
        logger.Error("Error could not get lastInsertedId for AddUserRequest", zap.Error(err))
        return nil, err
    }

	response :=  &fs.AddUserResponse{
        Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

        Count: 1,
        RecordId: cast.ToString(insertedId),
    }

    return response, nil
}

func (se *Executor) ExecuteAddUserBulk(ctx context.Context, bulkRequest *fs.BulkAddUserRequest) (*fs.BulkAddUserResponse, error) {

	var args []interface{}
	currQuery := database.QUERY_ADDUSER
    if idx := strings.Index(currQuery, "(?"); idx != -1 {
        currQuery = currQuery[:idx]
    }

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			currQuery += "(?,?);"
		} else {
			currQuery += "(?,?);,"
		}
		model := mappers.MakeAddUserRequestVO(request)
		args = append(args, AddUserArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
    if err != nil {
        logger.Error("Error could not get lastInsertedId for BulkAddUserRequest", zap.Error(err))
        return nil, err
    }

	var responses []*fs.AddUserResponse
	for index := range bulkRequest.Requests {
        responses = append(responses, &fs.AddUserResponse{
            Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

            Count: 1,
            RecordId: cast.ToString(cast.ToInt(insertedId)+index),
        })
    }

    response := &fs.BulkAddUserResponse{
        Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

        Count: cast.ToInt32(len(bulkRequest.Requests)),
        Responses: responses,
    }

	return response, nil
}

func (se *Executor) ExecuteFindBlood(ctx context.Context, request *fs.FindBloodRequest) (*fs.FindBloodResponse , error) {

	response := &fs.FindBloodResponse{}
	var rows = entsql.Rows{}
	model := mappers.MakeFindBloodRequestVO(request)
    args := FindBloodArgsReq(model)
	currQuery := database.QUERY_FINDBLOOD
	
	err := Driver.GetDriver().Query(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindBloodRequest", zap.Error(err))
        return nil, err
	}
	for rows.Next() {
		model := models.FindBloodResponseVO{}
		err := rows.Scan(&model.Name,&model.Location,&model.BloodType,&model.Gender,&model.PhoneNumber)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindBloodRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindBloodResponseVO(&model))
	}
	response.Status = &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        }
	return response, nil
}

func (se *Executor) ExecuteAddBlood(ctx context.Context, request *fs.AddBloodRequest) (*fs.AddBloodResponse, error) {

	model := mappers.MakeAddBloodRequestVO(request)
	args := AddBloodArgs(model)

	var rows sql.Result
	currQuery := database.QUERY_ADDBLOOD
    
	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddBloodRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
    if err != nil {
        logger.Error("Error could not get lastInsertedId for AddBloodRequest", zap.Error(err))
        return nil, err
    }

	response :=  &fs.AddBloodResponse{
        Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

        Count: 1,
        RecordId: cast.ToString(insertedId),
    }

    return response, nil
}

func (se *Executor) ExecuteAddBloodBulk(ctx context.Context, bulkRequest *fs.BulkAddBloodRequest) (*fs.BulkAddBloodResponse, error) {

	var args []interface{}
	currQuery := database.QUERY_ADDBLOOD
    if idx := strings.Index(currQuery, "(?"); idx != -1 {
        currQuery = currQuery[:idx]
    }

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			currQuery += "(?,?,?,?,?);"
		} else {
			currQuery += "(?,?,?,?,?);,"
		}
		model := mappers.MakeAddBloodRequestVO(request)
		args = append(args, AddBloodArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddBloodRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
    if err != nil {
        logger.Error("Error could not get lastInsertedId for BulkAddBloodRequest", zap.Error(err))
        return nil, err
    }

	var responses []*fs.AddBloodResponse
	for index := range bulkRequest.Requests {
        responses = append(responses, &fs.AddBloodResponse{
            Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

            Count: 1,
            RecordId: cast.ToString(cast.ToInt(insertedId)+index),
        })
    }

    response := &fs.BulkAddBloodResponse{
        Status: &fs.Status{
            Status: fs.StatusCode_SUCCESS,
        },

        Count: cast.ToInt32(len(bulkRequest.Requests)),
        Responses: responses,
    }

	return response, nil
}




func Execute(ctx context.Context, request *fs.MultiRequests) []error {
/**
	var response []error
	var err error
	multiRequest := request.Request
	for _, customerRequest := range multiRequest {
		addRequest := customerRequest.ARequest
		switch addRequest.(type) {
		case *fs.Request_ReqAddCustomer:
			modifiedRequest := addRequest.(*fs.Request_ReqAddCustomer)
			err = ExecuteAddCustomer(ctx, modifiedRequest.ReqAddCustomer)
			break
		case *fs.Request_ReqAddCustomerBulk:
			modifiedRequest := addRequest.(*fs.Request_ReqAddCustomerBulk)
			err = ExecuteAddCustomerBulk(ctx, modifiedRequest.ReqAddCustomerBulk)
			break
		default:
			logger.Info("Unkown request type")
			break
		}
		response = append(response, err)
	}
	return response
 */
	return nil;
}

