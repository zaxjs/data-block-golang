package data_block

type BaseResponseModel[T interface{}] struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Success bool   `json:"success,omitempty"`
}
