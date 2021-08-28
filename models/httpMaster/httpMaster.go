package httpMaster

func HttpRequest(surl string, acticon string, args map[string]interface{}, username string, ip string) (ret string, err error) {
	return WebDataBaseOpr(surl, acticon, args, username, ip)
}
