package request

type APIRequest[T any] struct {
	Data T `json:"data"`
}
