package executor

import (
    fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/models"
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

func FindPasswordArgs(request *fs.FindPasswordRequest) []interface{} {

	var args []interface{}
	args = append(args, request.Id)

	return args
}

func FindPasswordArgsReq(model *models.FindPasswordRequestVO) ([]interface{}){

	var args []interface{}
	args = append(args, model.Id)

	return args
}
func AddUserArgs(model *models.AddUserRequestVO) []interface{}{

	var args []interface{}
	args = append(args, model.Id)
	args = append(args, model.Password)

	return args
}
func FindBloodArgs(request *fs.FindBloodRequest) []interface{} {

	var args []interface{}
	args = append(args, request.BloodType)
	args = append(args, request.Location)

	return args
}

func FindBloodArgsReq(model *models.FindBloodRequestVO) ([]interface{}){

	var args []interface{}
	args = append(args, model.BloodType)
	args = append(args, model.Location)

	return args
}
func AddBloodArgs(model *models.AddBloodRequestVO) []interface{}{

	var args []interface{}
	args = append(args, model.Name)
	args = append(args, model.Location)
	args = append(args, model.BloodType)
	args = append(args, model.Gender)
	args = append(args, model.PhoneNumber)

	return args
}

