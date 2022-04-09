package model

type Status struct {
	Message string `json:"message"`
	Status  int32  `json:"status"`
}

type JwtRes struct {
	Message string `json:"message"`
	Status  int32  `json:"status"`
	Token   string `json:"token"`
}
