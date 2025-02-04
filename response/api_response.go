package response

type APIResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
