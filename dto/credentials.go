package dto

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWT struct {
	ErrorMsg *string `json:"error_msg"`
	Data     struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	} `json:"data"`
}
