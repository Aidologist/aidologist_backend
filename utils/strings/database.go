package strings

func MysqlConnectionString(
	username string,
	password string,
	registerDataBaseName string,
	charset string,
	url string,
	port string) string {
	return username+":"+password+"@tcp("+url+":"+port+")/"+registerDataBaseName+"?charset="+charset
}