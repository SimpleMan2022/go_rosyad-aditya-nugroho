package errorHandler_packages_jwt

type ErrorHandler struct {
	Status     bool   `json: status`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
