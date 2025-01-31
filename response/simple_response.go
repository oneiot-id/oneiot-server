package response

type SimpleBatchResponse struct {
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

type SimpleResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SimpleErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
