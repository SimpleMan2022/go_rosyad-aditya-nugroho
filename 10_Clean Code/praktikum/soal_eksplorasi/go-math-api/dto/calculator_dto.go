package dto

type CalculatorRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type CalculatorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     int    `json:"result"`
}
