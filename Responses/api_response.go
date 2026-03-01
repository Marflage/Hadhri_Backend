package responses

type ApiResponse[T any] struct {
	// TODO: Use the omitempty validation tag where appropriate.
	Error string `json:"error"`
	// TODO: How to make data nullable/optional here?
	Data    T      `json:"data"`
	Message string `json:"message"`
}
