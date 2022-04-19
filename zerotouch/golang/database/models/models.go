package models

import "database/sql"

type FindPasswordResponseVO struct {
		Password			sql.NullString
}
type FindPasswordRequestVO struct {
		Id			sql.NullString
}
type AddUserRequestVO struct {
		Id			sql.NullString
		Password			sql.NullString
}
type FindBloodResponseVO struct {
		Name			sql.NullString
		Location			sql.NullString
		BloodType			sql.NullString
		Gender			sql.NullString
		PhoneNumber			sql.NullString
}
type FindBloodRequestVO struct {
		BloodType			sql.NullString
		Location			sql.NullString
}
type AddBloodRequestVO struct {
		Name			sql.NullString
		Location			sql.NullString
		BloodType			sql.NullString
		Gender			sql.NullString
		PhoneNumber			sql.NullString
}

