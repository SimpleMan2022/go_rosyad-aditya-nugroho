package errorHandler_fruits

type ErrorHandler struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
