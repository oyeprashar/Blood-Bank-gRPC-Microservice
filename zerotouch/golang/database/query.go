package database

const (
	QUERY_FINDPASSWORD = "SELECT password FROM userLogins WHERE id = ?;"
	QUERY_ADDUSER = "INSERT INTO userLogins VALUES (?,?);"
	QUERY_FINDBLOOD = "SELECT * FROM bloodDB WHERE blood_type = ? AND location = ?;"
	QUERY_ADDBLOOD = "INSERT INTO bloodDB VALUES (?,?,?,?,?);"
	
)
