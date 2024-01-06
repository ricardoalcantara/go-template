package domain

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error   string         `json:"error"`
	Details []ErrorDetail `json:"details,omitempty"`
}
