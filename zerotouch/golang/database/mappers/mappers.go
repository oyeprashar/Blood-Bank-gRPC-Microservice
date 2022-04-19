package mappers

import (
	"code.nurture.farm/BloodBankSystemService/zerotouch/golang/database/models"
	fs "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
	"github.com/golang/protobuf/ptypes"
	"database/sql"
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeFindPasswordResponseVO(model *models.FindPasswordResponseVO) *fs.FindPasswordResponseRecord {

	return &fs.FindPasswordResponseRecord { 
		Password: model.Password.String,
	}

}
func MakeFindPasswordRequestVO(request *fs.FindPasswordRequest) *models.FindPasswordRequestVO {

	return &models.FindPasswordRequestVO { 
		Id: getNullableString(request.Id),
	}

}
func MakeAddUserRequestVO(request *fs.AddUserRequest) *models.AddUserRequestVO {

	return &models.AddUserRequestVO { 
		Id: getNullableString(request.Id),
		Password: getNullableString(request.Password),
	}

}
func MakeFindBloodResponseVO(model *models.FindBloodResponseVO) *fs.FindBloodResponseRecord {

	return &fs.FindBloodResponseRecord { 
		Name: model.Name.String,
		Location: model.Location.String,
		BloodType: model.BloodType.String,
		Gender: model.Gender.String,
		PhoneNumber: model.PhoneNumber.String,
	}

}
func MakeFindBloodRequestVO(request *fs.FindBloodRequest) *models.FindBloodRequestVO {

	return &models.FindBloodRequestVO { 
		BloodType: getNullableString(request.BloodType),
		Location: getNullableString(request.Location),
	}

}
func MakeAddBloodRequestVO(request *fs.AddBloodRequest) *models.AddBloodRequestVO {

	return &models.AddBloodRequestVO { 
		Name: getNullableString(request.Name),
		Location: getNullableString(request.Location),
		BloodType: getNullableString(request.BloodType),
		Gender: getNullableString(request.Gender),
		PhoneNumber: getNullableString(request.PhoneNumber),
	}

}


func getNullableInt32(nullableInt int32) sql.NullInt32{

	var result sql.NullInt32
	if nullableInt != 0 {
		result = sql.NullInt32{Int32: nullableInt, Valid: true}
	} else {
		result = sql.NullInt32{}
	}
	return result
}

func getNullableInt32s(nullableInts []int32) []sql.NullInt32{

	var result []sql.NullInt32
	if len(nullableInts) > 0 {
		for _, nullableInt := range nullableInts {
			result = append(result, sql.NullInt32{Int32: nullableInt, Valid: true})
		}
	} else {
		result = []sql.NullInt32{}
	}
	return result
}

func getNullableInt64(nullableInt int64) sql.NullInt64{

	var result sql.NullInt64
	if nullableInt != 0 {
		result = sql.NullInt64{Int64: nullableInt, Valid: true}
	} else {
		result = sql.NullInt64{}
	}
	return result
}

func getNullableInt64s(nullableInts []int64) []sql.NullInt64{

	var result []sql.NullInt64
	if len(nullableInts) > 0 {
		for _, nullableInt := range nullableInts {
			result = append(result, sql.NullInt64{Int64: nullableInt, Valid: true})
		}
	} else {
		result = []sql.NullInt64{}
	}
	return result
}

func getNullableFloat64(nullableFloat float64) sql.NullFloat64 {

	var result sql.NullFloat64
	if &nullableFloat != nil {
		result = sql.NullFloat64{Float64: nullableFloat, Valid: true}
	} else {
		result = sql.NullFloat64{}
	}
	return result
}

func getNullableString(nullableString string) sql.NullString{

	var result sql.NullString
	if len(nullableString) > 0 {
		result = sql.NullString{String: nullableString, Valid: true}
	} else {
		result = sql.NullString{}
	}
	return result
}

func getNullableStrings(nullableStrings []string) []sql.NullString{

	var result []sql.NullString
	if len(nullableStrings) > 0 {
		for _, nullableString := range nullableStrings {
			result = append(result, sql.NullString{String: nullableString, Valid: true})
		}
	} else {
		result = []sql.NullString{}
	}
	return result
}

func getNullableDateTime(timeStamp int64) sql.NullTime{

	gTime := time.Unix(timeStamp, 0)
	pTime, _ := ptypes.TimestampProto(gTime)

	var result sql.NullTime
	parsedAllottedTime, err := ptypes.Timestamp(pTime)
	if err != nil {
		result = sql.NullTime{}
	} else {
		result = sql.NullTime{Time: parsedAllottedTime, Valid: true}
	}
	return result
}

func getNullableTimestamp(timeStamp *timestamppb.Timestamp) sql.NullString{

	  //gTime := time.Unix(timeStamp, 0)
  	//pTime, _ := ptypes.TimestampProto(gTime)

  	var result sql.NullString
  	parsedAllottedTime, err := ptypes.Timestamp(timeStamp)
  	if err != nil {
  		result = sql.NullString{}
  	} else if parsedAllottedTime.IsZero() {
        	result = sql.NullString{}
      } else {
  		result = sql.NullString{String: parsedAllottedTime.Format("2006-01-02 15:04:05"), Valid: true}
  	}
  	return result
  }

func getNullableBool(boolean bool) sql.NullBool {

	var result sql.NullBool
	result = sql.NullBool{Bool: boolean, Valid: true}
	return result
}

func getUnixTime(timestamp string) int64{

	layout := "2006-01-02T15:04:05Z"
	parsedTime, _ := time.Parse(layout, timestamp)
	return parsedTime.Unix()
}

func getProtoTime(timestamp string) *timestamppb.Timestamp {
	layout := "2006-01-02T15:04:05Z"
	time, _ := time.Parse(layout, timestamp)
	return timestamppb.New(time)
}
