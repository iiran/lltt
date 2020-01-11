package core

type ArrayResponse struct {
	Data   interface{} `json:"data"`
	Remain int64       `json:"remain"`
}

type ObjectResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Detail string `json:"detail"`
}
