package types

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Error *ResponseError `json:"error,omitempty"`
}
