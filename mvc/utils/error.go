package utils

type AppError struct {
	ID      int    `json:"error_code"`
	Message string `json:"message"`
}
