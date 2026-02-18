package responses

type ApiResponse struct {
	// TODO: Use the omitempty validation tag where appropriate.
	Error   string `json:"error"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
